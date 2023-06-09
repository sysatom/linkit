"use client"

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {useEffect, useState} from "react";
import fetcher from "@/helpers/http";
import {OPERATE} from "@/constants/ACTION";
import Dropdown from "@/components/ui/dropdown";

export function BotInfo({ bots }) {
  let [help, setHelp] = useState<{ [key: string]: Array<String> }>({});
  let [selectedBot, setSelectedBot] = useState("")

  useEffect(() => {
    if (selectedBot != "") {
      fetcher(OPERATE.help, selectedBot).then(r => setHelp(r)).catch(console.error);
    }
  }, [selectedBot])

  let availableBots = []
  if (bots != null && bots.bots.length > 0) {
    bots.bots
      .map(bot => {
        availableBots.push({
          label: bot.name,
          id: bot.id,
        })
      });

    availableBots?.splice(0, 0, {
      label: "All Bots",
      id: "",
    });
  }

  const handleBotChange = (id: string) => {
    setSelectedBot(id);
  };

  return (
    <div>
      <Dropdown
        triggerLabel={
          availableBots?.find((project) => project.id === selectedBot)
            ?.label
        }
        items={availableBots}
        selectedId={selectedBot}
        handleOnClick={handleBotChange}
      />
      <Card>
        <CardHeader>
          <CardTitle>{selectedBot ?? "-"}</CardTitle>
          <CardDescription>Bots help info</CardDescription>
        </CardHeader>
        <CardContent className="grid gap-6">
          <table style={{width: "90%", margin: "auto"}}>
            <thead>
            <tr>
              <th>Type</th>
              <th>Help</th>
            </tr>
            </thead>
            <tbody>
            {
              help != null ? Object.keys(help).map((key, i) => {
                return help[key].map((v, j) => <tr key={`${i}-${j}`}>
                  <td className="font-bold">{key}</td>
                  <td className="text-left">{v}</td>
                </tr>)
              }) : <tr><td colSpan={2} className="text-center">empty</td></tr>
            }
            </tbody>
          </table>
        </CardContent>
      </Card>
    </div>
  )
}
