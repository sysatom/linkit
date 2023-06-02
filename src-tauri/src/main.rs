// Prevents additional console window on Windows in release
#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

mod agent;
mod util;
mod instruct;

use tauri::Manager;
use tokio::sync::{mpsc, Mutex};
use tracing::info;

struct AsyncProcInputTx {
    inner: Mutex<mpsc::Sender<String>>
}

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", name)
}

#[tauri::command]
async fn event_emit(message: String, state: tauri::State<'_, AsyncProcInputTx>) -> Result<(), String> {
    info!(?message, "event_emit");
    let async_proc_input_tx = state.inner.lock().await;
    async_proc_input_tx
        .send(message)
        .await
        .map_err(|e| e.to_string())
}

fn event_send_frontend<R: tauri::Runtime>(message: String, manager: &impl Manager<R>) {
    info!(?message, "event_send_frontend");
    manager
        .emit_all("event_broadcast", format!("{}", message))
        .unwrap();
}

fn main() {
    tracing_subscriber::fmt::init();

    let (async_proc_input_tx, async_proc_input_rx) = mpsc::channel(1);
    let (async_proc_output_tx, mut async_proc_output_rx) = mpsc::channel(1);

    tauri::Builder::default()
        .plugin(tauri_plugin_store::Builder::default().build())
        .manage(AsyncProcInputTx {
            inner: Mutex::new(async_proc_input_tx),
        })
        .invoke_handler(tauri::generate_handler![greet, event_emit])
        .setup(|app| {
            tauri::async_runtime::spawn(async move {
                async_process_model(
                    async_proc_input_rx,
                    async_proc_output_tx,
                ).await
            });

            let app_handle = app.handle();
            tauri::async_runtime::spawn(async move {
                loop {
                    if let Some(output) = async_proc_output_rx.recv().await {
                        event_send_frontend(output, &app_handle);
                    }
                }
            });

            Ok(())
        })
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}

async fn async_process_model(
    mut input_rx: mpsc::Receiver<String>,
    output_tx: mpsc::Sender<String>,
) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    loop {
        while let Some(input) = input_rx.recv().await {
            let output = input;
            output_tx.send(output).await?;
        }
    }
}
