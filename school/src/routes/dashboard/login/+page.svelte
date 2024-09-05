<script lang="ts">
    import { goto } from "$app/navigation";
    import { Label, Input, Button } from "flowbite-svelte";

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
        if (data.success) {
            console.log("Login successful");
            goto(`/dashboard/${schoolName}`);
        } else {
            console.log("Login failed");
        }
    }
</script>

<div class="mt-2 flex items-center justify-center p-4">
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
