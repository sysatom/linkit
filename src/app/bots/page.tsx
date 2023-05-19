import {BotsSwitch} from "@/app/bots/components/bots-switch";

export default function BotsPage() {
  return (
    <div className="items-start justify-center gap-6 rounded-lg p-8 md:grid lg:grid-cols-2 xl:grid-cols-3">
      <div className="col-span-2 grid items-start gap-6 lg:col-span-1">
      <BotsSwitch />
      </div>
    </div>
  )
}
