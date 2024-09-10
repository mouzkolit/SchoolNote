<script lang="ts">
    import {
        Card,
        Button,
        Input,
        Label,
        Breadcrumb,
        BreadcrumbItem,
    } from "flowbite-svelte";

    // Placeholder data for classes (replace with actual data fetching logic)
    let classes = [
        { id: 1, name: "Mathematics", teacher: "John Doe", students: 25 },
        { id: 2, name: "Science", teacher: "Jane Smith", students: 30 },
        // Add more classes as needed
    ];

    let newClassName = "";
    let newTeacherName = "";

    function addClass() {
        // Add logic to create a new class
        console.log("Adding class:", newClassName, newTeacherName);
        // Reset input fields
        newClassName = "";
        newTeacherName = "";
    }
</script>

<div class="m-4">
    <Breadcrumb>
        <BreadcrumbItem href="/">Home</BreadcrumbItem>
        <BreadcrumbItem href="/dashboard">Dashboard</BreadcrumbItem>
        <BreadcrumbItem href="/dashboard/class" current>Class</BreadcrumbItem>
    </Breadcrumb>
    <div class="m-4">
        <Card class="w-screen flex flex-grow" size="xl">
            <Label>Class Name</Label>
            <Input class="my-4" />
            <Label>Teacher</Label>
            <Input class="my-4" />
            <Label>Class Name</Label>
            <Input class="my-4" />
            <Button class="my-4 bg-blue-400">Create Class</Button>
        </Card>
    </div>
</div>
