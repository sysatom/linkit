export interface IUser {
    user: {
        id: string;
        email: string;
        name: string;
        username: string;
        avatar: string;
    };
}

export interface IBot {
    bots: {
        id: string;
        name: string;
    }[];
}
