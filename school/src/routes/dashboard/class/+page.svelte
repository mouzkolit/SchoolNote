<script lang="ts">
    import {
        Card,
        Button,
        Input,
        Label,
        Breadcrumb,
        BreadcrumbItem,
    } from "flowbite-svelte";
    import { schoolStore } from "$lib/schoolStore";

    let currentSchoolId: string = "";
    schoolStore.subscribe((value) => {
        currentSchoolId = value.schoolId;
    })();
    // Placeholder data for classes (replace with actual data fetching logic)
    let className: string = "";
    let teacherFirst: string = "";
    let teacherLast: string = "";

    async function addClass() {
        // Add logic to create a new class
        if (className === "" || teacherFirst === "" || teacherLast === "") {
            alert("Please enter all the necessary information for your class");
        }

        const response = await fetch(
            `http://localhost:8080/${currentSchoolId}/classes`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({
                    className: className,
                    teacherFirst: teacherFirst,
                    teacherLast: teacherLast,
                }),
            },
        );

        if (response.ok) {
            const data = await response.json();
            console.log("Class created successfully:", data);
        } else {
            console.error("Failed to create class");
        }
    }
</script>

<div class="m-4">
    <Breadcrumb>
        <BreadcrumbItem href="/">Home</BreadcrumbItem>
        <BreadcrumbItem href="/class" current>Class</BreadcrumbItem>
    </Breadcrumb>
    <div class="m-4">
        <Card class="w-screen flex flex-grow" size="xl">
            <Label>Class Name</Label>
            <Input class="my-4" bind:value={className} />
            <Label>Teacher First Name</Label>
            <Input class="my-4" bind:value={teacherFirst} />
            <Label>Teacher Last Name</Label>
            <Input class="my-4" bind:value={teacherLast} />
            <Button class="my-4 bg-blue-400" on:click={addClass}
                >Create Class</Button
            >
        </Card>

        <Card class="w-screen flex flex-grow"></Card>
    </div>
</div>
