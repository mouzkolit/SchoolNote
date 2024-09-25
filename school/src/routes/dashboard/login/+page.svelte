<script lang="ts">
    import { goto } from "$app/navigation";
    import { Label, Input, Button } from "flowbite-svelte";
    import { isAuthenticated, currentSchool } from "$lib/authStore";

    let schoolName: string | null = null;
    let password: string | null = null;

    async function login() {
        if (!schoolName || !password) {
            console.error("School name and password are required");
            return;
        }
        var url = new URL("http://localhost:8080/school/login");
        var params = {
            name: schoolName,
            password: password,
        };
        url.search = new URLSearchParams(params).toString();
        const response = await fetch(url, {
            method: "POST",
            credentials: "include",
        });
        const data = await response.json();
        console.log(data);
        if (data.message == "Login successful") {
            console.log("Login successful");
            isAuthenticated.set(true);
            currentSchool.set(schoolName);
            //goto(`/dashboard/${schoolName}`);
        } else {
            console.log("Login failed");
        }
    }
</script>

<div class="mt-8 flex items-center justify-center h-screen">
    <div class="bg-gray-200 p-8 rounded shadow-lg w-full max-w-md">
        <form>
            <div class="mb-6">
                <Label for="school-name" class="block mb-2">School Name</Label>
                <Input
                    id="school-name"
                    placeholder="Enter your school name"
                    class="border p-2 w-full"
                    bind:value={schoolName}
                />
            </div>
            <div class="mb-6">
                <Label for="password" class="block mb-2">Password</Label>
                <Input
                    id="password"
                    type="password"
                    placeholder="Enter your password"
                    class="border p-2 w-full"
                    bind:value={password}
                />
            </div>
            <Button
                type="submit"
                class="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                on:click={login}
            >
                Login
            </Button>
        </form>
    </div>
</div>
