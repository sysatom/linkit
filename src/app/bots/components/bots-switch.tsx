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
import {useEffect, useState} from "react";
import fetcher from "@/helpers/http";
import {OPERATE} from "@/constants/ACTION";
import {IBot} from "@/types";

export function BotsSwitch() {
  const [bots, setBots] = useState<IBot>(null)

  useEffect(()=> {
    fetcher(OPERATE.bots).then(v => setBots(v)).catch(console.error)
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
              These cookies are essential in order to use the website and use
              its features.
            </span>
            </Label>
            <Switch id={i.id} defaultChecked />
          </div>) :
            <div className="text-center">empty</div>
        }
      </CardContent>
    </Card>
  )
}
