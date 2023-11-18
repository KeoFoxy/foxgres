export interface LoginProps {
    email: string;
    password: string;
}

export const fetchLogin = async ({ email, password }: LoginProps) => {
    try {
        const response = await fetch(`URL_API/login`, {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        });

        if (response.ok) {
            const data = await response.json();
            return data;
        } else {
            return undefined; // TODO: think of it
        }
    } catch (error) {
        throw error;
    }
};
