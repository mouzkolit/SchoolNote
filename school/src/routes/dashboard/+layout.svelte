<script lang="ts">
    import "../app.css";
    import { Label, Input, Button } from "flowbite-svelte";
</script>

<div class="mt-20 flex items-center justify-center min-h-screen">
    <div class="bg-gray-200 p-8 rounded shadow-lg w-full max-w-md">
        <form>
            <div class="mb-6">
                <Label for="school-name" class="block mb-2">School Name</Label>
                <Input
                    id="school-name"
                    placeholder="Enter your school name"
                    class="border p-2 w-full"
                />
            </div>
            <div class="mb-6">
                <Label for="password" class="block mb-2">Password</Label>
                <Input
                    id="password"
                    type="password"
                    placeholder="Enter your password"
                    class="border p-2 w-full"
                />
            </div>
            <Button
                type="submit"
                class="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
                Login
            </Button>
        </form>
    </div>
</div>
