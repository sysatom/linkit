"use client"

import {BotsSwitch} from "@/app/bots/components/bots-switch";
import {BotInfo} from "@/app/bots/components/bot-info";
import {useEffect, useState} from "react/index";
import {IBot} from "@/types";
import fetcher from "@/helpers/http";
import {OPERATE} from "@/constants/ACTION";
import "./styles.css"

export default function BotsPage() {
  const [bots, setBots] = useState<IBot>(null)

  useEffect(()=> {
    fetcher(OPERATE.bots).then(v => setBots(v)).catch(console.error)
  }, [])

  return (
    <div className="items-start justify-center gap-6 rounded-lg p-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <BotsSwitch bots={bots} />
      </div>
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <BotInfo bots={bots} />
      </div>
    </div>
  )
}
