import { API_BASE_URL } from "../config";

export interface LoginProps {
  login: string;
  password: string;
}

export const fetchLogin = async ({ login, password }: LoginProps) => {
  const response = await fetch(`${API_BASE_URL}/auth`, {
    method: "GET",
    headers: {
      "Content-type": "application/json",
      login: login,
      password: password,
    },
  });

  if (response.ok) {
    const data = await response.json();
    return data;
  } else {
    return undefined; // TODO: think of it
  }
};
