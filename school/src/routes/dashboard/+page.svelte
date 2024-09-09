<script lang="ts">
    import { goto } from "$app/navigation";
    import { Button, Card, Heading, P } from "flowbite-svelte";

    function startLogin() {
        goto("/dashboard/login");
    }
</script>

<div class="flex justify-center items-center h-screen bg-gray-100">
    <Card class="w-full max-w-3xl">
        <div class="flex items-center mb-4">
            <Heading tag="h2" class="text-3xl font-bold text-gray-900">
                Welcome to the Admin Panel
            </Heading>
        </div>
        <P class="mb-4 text-gray-700">
            Manage your school's information, students, and resources all in one
            place. This powerful dashboard provides you with the tools you need
            to efficiently administer your educational institution.
        </P>
        <P class="mb-6 text-gray-700">
            To get started, please log in with your school's credentials.
        </P>
        <div class="flex justify-center">
            <Button gradient color="blue" on:click={startLogin}>Log In</Button>
        </div>
    </Card>
</div>
