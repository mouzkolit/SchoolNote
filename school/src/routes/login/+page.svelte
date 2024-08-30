<script lang="ts">
    import { Section, Register } from "flowbite-svelte-blocks";
    import { Button, Checkbox, Label, Input } from "flowbite-svelte";
    import { LanguageOutline } from "flowbite-svelte-icons";

    let schoolName: string = "";
    let password: string = "";

    async function login() {
        var url = new URL("http://localhost:8080/school/login");
        var params = {
            name: schoolName,
            password: password,
        };
        url.search = new URLSearchParams(params).toString();
        const response = await fetch(url, {
            method: "POST",
        });
    }
</script>

<Section name="login">
    <Register href="/">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
            <form class="flex flex-col space-y-6" action="/">
                <h3
                    class="text-xl font-medium text-gray-900 dark:text-white p-0"
                >
                    Change Password
                </h3>
                <Label class="space-y-2">
                    <span>School Name</span>
                    <Input
                        type="text"
                        name="schoolName"
                        placeholder="enter"
                        required
                        bind:value={schoolName}
                    />
                </Label>
                <Label class="space-y-2">
                    <span>Your password</span>
                    <Input
                        type="password"
                        name="password"
                        placeholder="•••••"
                        required
                        bind:value={password}
                    />
                </Label>
                <div class="flex items-start">
                    <Checkbox>Remember me</Checkbox>
                    <a
                        href="/"
                        class="ml-auto text-sm text-blue-700 hover:underline dark:text-blue-500"
                        >Forgot password?</a
                    >
                </div>
                <Button type="submit" class="w-full1" on:click={login}
                    >Sign in</Button
                >
                <p class="text-sm font-light text-gray-500 dark:text-gray-400">
                    Don’t have a School Account yet? <a
                        href="/school"
                        class="font-medium text-primary-600 hover:underline dark:text-primary-500"
                        >Sign up</a
                    >
                </p>
            </form>
        </div>
    </Register>
</Section>
