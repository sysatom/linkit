"use client"

import { useEffect, useState } from "react"
import type { WebviewWindow } from "@tauri-apps/api/window"
import { Globe, Maximize, Mic, X } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  Menubar,
  MenubarCheckboxItem,
  MenubarContent,
  MenubarItem,
  MenubarLabel,
  MenubarMenu,
  MenubarRadioGroup,
  MenubarRadioItem,
  MenubarSeparator,
  MenubarShortcut,
  MenubarSub,
  MenubarSubContent,
  MenubarSubTrigger,
  MenubarTrigger,
} from "@/components/ui/menubar"

import { ExamplesNav } from "./examples-nav"
import { Icons } from "./icons"
import { ModeToggle } from "./mode-toggle"

export function Menu() {
  const [appWindow, setAppWindow] = useState<null | WebviewWindow>(null)

  // Dinamically import the tauri API, but only when it's in a tauri window
  useEffect(() => {
    import("@tauri-apps/api/window").then(({ appWindow }) => {
      setAppWindow(appWindow)
    })
  }, [])

  const minimizeWindow = () => {
    appWindow?.minimize()
  }
  const maximizeWindow = async () => {
    if (await appWindow?.isMaximized()) {
      appWindow?.unmaximize()
    } else {
      appWindow?.maximize()
    }
  }
  const closeWindow = () => {
    appWindow?.close()
  }

  const openPreferences = () => {
    // router.push("/preferences")
  }

  return (
    <Menubar className="rounded-none border-b border-none pl-2 lg:pl-4">
      <MenubarMenu>
        <MenubarTrigger className="font-bold">App</MenubarTrigger>
        <MenubarContent>
          <MenubarItem>About</MenubarItem>
          <MenubarSeparator />
          <a href="/preferences">
            <MenubarItem onClick={openPreferences}>
              Preferences... <MenubarShortcut>⌘,</MenubarShortcut>
            </MenubarItem>
          </a>
          <MenubarSeparator />
          <MenubarItem onClick={closeWindow}>
            Quit <MenubarShortcut>⌘Q</MenubarShortcut>
          </MenubarItem>
        </MenubarContent>
      </MenubarMenu>
      <MenubarMenu>
        <a href="/dashboard" className="text-sm shadow mr-4">Dashboard</a>
      </MenubarMenu>
      <MenubarMenu>
        <a href="/bots" className="text-sm shadow mr-4">Bots</a>
      </MenubarMenu>
      <div
        data-tauri-drag-region
        className="inline-flex h-full w-full justify-end"
      >
        <ExamplesNav />
        <div className="pr-3">
          <ModeToggle />
        </div>

        <Button
          onClick={minimizeWindow}
          variant="ghost"
          className="h-8 focus:outline-none"
        >
          <Icons.minimize className="h-3 w-3" />
        </Button>
        <Button
          onClick={maximizeWindow}
          variant="ghost"
          className="h-8 focus:outline-none hidden"
        >
          <Maximize className="h-4 w-4" />
        </Button>
        <Button
          onClick={closeWindow}
          variant="ghost"
          className="h-8 focus:outline-none"
        >
          <X className="h-4 w-4" />
        </Button>
      </div>
    </Menubar>
  )
}
