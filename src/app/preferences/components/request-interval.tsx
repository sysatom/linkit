import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card"
import {Input} from "@/components/ui/input";

export function RequestInterval() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Request Interval</CardTitle>
        <CardDescription>
          How frequently should Linkit check for new action?
        </CardDescription>
      </CardHeader>
      <CardContent className="pt-6">
        <div className="space-y-2">
          <Input id="password" type="text" placeholder="60 sec" />
        </div>
      </CardContent>
    </Card>
  )
}
