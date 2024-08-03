"use client";
/*
import { useState } from 'react'
import { Button } from "@/components/ui/button"

interface User {
    id: number;
    name: string;
    lastname: string;
    codpucp: string;
}

export default function Home() {
    const [users, setUsers] = useState<User[]>([])

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:1323/data')
            const data: User[] = await response.json()
            setUsers(data)
        } catch (error) {
            console.error("Error fetching data:", error)
        }
    }

    return (
        <div>
            <Button onClick={fetchData}>Click me</Button>
            <ul>
                {users.map(user => (
                    <li key={user.id}>
                        {user.name} {user.lastname} - {user.codpucp}
                    </li>
                ))}
            </ul>
        </div>
    )
}
*/
import { Button } from "@/components/ui/button";

export default function Home() {
	    const handleLogin = () => {
			        window.location.href = "http://localhost:1323/auth/google";
							    };

	    return (
        <div>
            <Button onClick={handleLogin}>Login with Google</Button>
        </div>
    );
}
