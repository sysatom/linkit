pub mod dev;
pub mod clipboard;
pub mod anki;

use anyhow::{anyhow, Result};

pub const AGENT_VERSION: i32 = 1;
pub const URI: &str = "http://127.0.0.1:6060/extra/linkit";//todo

pub async fn agent_post_data(agent_uri: String, data: serde_json::Value) -> Result<bool> {
    let client = reqwest::Client::new();
    let response = client.post(agent_uri)
        .body(data.to_string())
        .send()
        .await?;

    // info!("agent_post_data request status {}", response.status());

    if response.status().is_success() {
        Ok(true)
    } else {
        Err(anyhow!("error agent post data"))
    }
}