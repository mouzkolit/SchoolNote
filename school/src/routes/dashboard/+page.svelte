<script lang="ts">
    import { goto } from "$app/navigation";
    import { Button, Card, Heading, P } from "flowbite-svelte";

    function startLogin() {
        goto("/dashboard/login");
    }
</script>

<div class="p-4 w-full h-full">
    <div class="grid grid-cols-1 gap-4 w-full">
        <div class="w-full">
            <Card class="w-full" padding="xl">
                <div class="flex items-center mb-4">
                    <Heading tag="h2" class="text-3xl font-bold text-gray-900">
                        Welcome to the Admin Panel
                    </Heading>
                </div>
                <P class="mb-4 text-gray-700">
                    Manage your school's information, students, and resources
                    all in one place. This powerful dashboard provides you with
                    the tools you need to efficiently administer your
                    educational institution.
                </P>
                <P class="mb-6 text-gray-700">
                    To get started, please log in with your school's
                    credentials.
                </P>
                <div class="flex justify-center">
                    <Button gradient color="blue" on:click={startLogin}
                        >Log In</Button
                    >
                </div>
            </Card>
        </div>
        <div class="w-full">
            <Card class="w-full" padding="xl">
                <p>Second Card Content</p>
            </Card>
        </div>
        <div class="w-full">
            <Card class="w-full" padding="xl">
                <p>Third Card Content</p>
            </Card>
        </div>
    </div>
</div>

<style>
    :global(.card) {
        width: 100% !important;
        max-width: none !important;
    }
</style>
