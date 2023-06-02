"use client"

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"
import {store} from "@/helpers/store";
import {useEffect, useState} from "react/index";
import { invoke } from '@tauri-apps/api/tauri';

export function BotsSwitch({ bots }) {
  let [botSwitch, setBotSwitch] = useState<Map<String, boolean>>(new Map<string, any>());
  const handleUpdateValue = (key:string, newValue:boolean) => {
    const newMap = new Map<String, boolean>(Array.from(botSwitch).map(([k, v]) => (k === key ? [k, newValue] : [k, v])));
    newMap.set(key, newValue);
    store.set("bot-switch", JSON.stringify(Array.from(newMap))).then(console.log).catch(console.error);
    setBotSwitch(newMap);
    // emit bot switch
    invoke('bot_switch', { message: JSON.stringify(Array.from(newMap)) })
      .then(console.log).catch(console.error)
  };

  useEffect(() => {
    const getBotSwitch = async () => {
      const value = await store.get("bot-switch");
      if (value) {
        let result = JSON.parse(value as string);
        if (result) {
          let m = result as Array<any>
          if (m != undefined) {
            const newMap = new Map<String, boolean>()
            m.forEach(([key, value]) => {
              newMap.set(key, value);
            });
            setBotSwitch(newMap);
          }
        }
      }
    }
    getBotSwitch().then(console.log).catch(console.error);
  }, [])

  return (
    <Card>
      <CardHeader>
        <CardTitle>Bots Instruct Settings</CardTitle>
        <CardDescription>Manage your Bots Instruct settings here.</CardDescription>
      </CardHeader>
      <CardContent className="grid gap-6">
        {
          bots != null && bots.bots.length > 0 ?
            bots.bots.map(i => <div key={i.id} className="flex items-center justify-between space-x-2">
            <Label htmlFor="necessary" className="flex flex-col space-y-1">
              <span>{i.name}</span>
              <span className="font-normal leading-snug text-muted-foreground">
              Bot detail ......
            </span>
            </Label>
            <Switch id={i.id}
                    defaultChecked={botSwitch.get(i.id)}
                    onCheckedChange={checked => {handleUpdateValue(i.id, checked)}} />
          </div>) :
            <div className="text-center">empty</div>
        }
      </CardContent>
    </Card>
  )
}
