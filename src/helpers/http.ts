import {store} from "@/helpers/store";
import {API_ENDPOINTS} from "@/constants/API";

export default async function fetcher<JSON = any>(
    action: string,
    content?: Object
): Promise<JSON> {
    const bearerToken = await store.get("token");

    if (!bearerToken) {
        throw new Error("No token found");
    }
    const data = {
        "action": action,
        "version": 1,
        "content": content
    }
    const res = await fetch(API_ENDPOINTS.base, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${bearerToken}`
        },
        body: JSON.stringify(data)
    });

    if (!res.ok) {
        throw new Error("Not authenticated");
    }

    return res.json();
}
