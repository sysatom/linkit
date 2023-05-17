import { Metadata } from "next"

import { cn } from "@/lib/utils"

import { OtherSettings } from "./components/other-settings"
import { AccessToken } from "./components/access-token"
import { RequestInterval } from "./components/request-interval"
import { Notifications } from "./components/notifications"
import { Logging } from "@/app/preferences/components/logging";
import "./styles.css"

export const metadata: Metadata = {
  title: "Cards",
  description: "Examples of cards built using the components.",
}

function Container({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn(
        "flex items-center justify-center [&>div]:w-full",
        className
      )}
      {...props}
    />
  )
}

export default function PreferencesPage() {
  return (
    <div className="items-start justify-center gap-6 rounded-lg p-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <Container>
          <AccessToken />
        </Container>
        <Container>
          <OtherSettings />
        </Container>
      </div>
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
        <Container>
          <RequestInterval />
        </Container>
        <Container>
          <Notifications />
        </Container>
        <Container>
          <Logging />
        </Container>
      </div>
    </div>
  )
}
