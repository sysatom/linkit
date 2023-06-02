"use client"

import {BotsSwitch} from "@/app/bots/components/bots-switch";
import {BotInfo} from "@/app/bots/components/bot-info";
import {useEffect, useState} from "react/index";
import {IBot} from "@/types";
import fetcher from "@/helpers/http";
import {OPERATE} from "@/constants/ACTION";
import "./styles.css"
import { listen } from '@tauri-apps/api/event';
import { invoke } from '@tauri-apps/api/tauri';

export default function BotsPage() {
  const [bots, setBots] = useState<IBot>(null)

  const [output, setOutput] = useState("")
  const sendOutput = () => {
    invoke('event_emit', { message: output })
      .then(console.log).catch(console.error)
  }

  useEffect(()=> {
    fetcher(OPERATE.bots).then(v => setBots(v)).catch(console.error)

    listen('event_broadcast', (event) => {
      console.log("event_broadcast: " + event.payload)
    }).then(console.log).catch(console.error)
  }, [])

  const handleChange = (event)=> {
      setOutput(event.target.value)
  }

  return (
    <div className="items-start justify-center gap-6 rounded-lg p-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <BotsSwitch bots={bots} />
      </div>
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <BotInfo bots={bots} />
      </div>
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <input type="text" value={output} onChange={handleChange} />
          <button onClick={sendOutput}>
            test
          </button>
      </div>
    </div>
  )
}
