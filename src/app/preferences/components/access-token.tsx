"use client"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {store} from "@/helpers/store";
import {useEffect, useState} from "react";

export function AccessToken() {
  const [token, setToken] = useState("")
  useEffect(()=> {
    store.get("token").then(v => setToken(String(v)))
  }, [])

  const saveToken = () => {
    if (token == "") {
      alert("please input token")
      return
    }
    store.set("token", token).then(console.log).catch(console.error)
  }

  return (
    <Card>
      <CardHeader className="space-y-1">
        <CardTitle className="text-2xl">Bot Access Token</CardTitle>
        <CardDescription>
          Enter your bot access token, Steps for creating a Access Token: Enter the command `/access token` in Linkit bot.
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-4">
        <div className="grid gap-2">
          <Label htmlFor="password">Token</Label>
          <Input id="password" type="text" value={token} onChange={e => setToken(e.target.value)}/>
        </div>
      </CardContent>
      <CardFooter>
        <Button className="w-full" onClick={saveToken}>Save</Button>
      </CardFooter>
    </Card>
  )
}
