<script>
    import { page } from "$app/stores";
    import { isAuthenticated, currentSchool } from "$lib/authStore";
    import { schoolStore } from "$lib/schoolStore";

    import {
        Sidebar,
        SidebarGroup,
        SidebarItem,
        SidebarWrapper,
        Heading,
        Span,
        Label,
        Button,
    } from "flowbite-svelte";
    import {
        ChartPieSolid,
        ArrowRightToBracketOutline,
        HomeSolid,
        BarsOutline,
    } from "flowbite-svelte-icons";
    $: activeUrl = $page.url.pathname;
    $: schoolName = $schoolStore.schoolName;
    $: schoolId = $schoolStore.schoolId;

    const toggleSidebar = () => (isOpen = !isOpen);

    // Add this line to export the isOpen variable
    export let isOpen = false;
</script>

<div class="flex">
    <Sidebar
        {activeUrl}
        class="transition-transform duration-300 ease-in-out {isOpen
            ? 'translate-x-0'
            : '-translate-x-full'}"
    >
        <SidebarWrapper
            class="bg-blue-200 h-screen fixed left-0 top-0 flex-col"
        >
            <Button
                class="absolute top-4 right-4 transition-opacity duration-300 ease-in-out bg-blue-300 "
                aria-label="Toggle sidebar"
                on:click={toggleSidebar}
            >
                <BarsOutline class="w-5 h-5" />
            </Button>
            <div>
                <Heading
                    tag="h1"
                    class="mb-4 p-4"
                    customSize="text-3xl font-extrabold md:text-5xl lg:text-6xl"
                >
                    <Span gradient>Admin</Span>.
                </Heading>
                <Label>School Name: {schoolName}</Label>
                <SidebarGroup>
                    {#if !$isAuthenticated}
                        <SidebarItem label="Login" href="/dashboard/login">
                            <svelte:fragment slot="icon">
                                <ArrowRightToBracketOutline
                                    class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                                />
                            </svelte:fragment>
                        </SidebarItem>
                    {:else}
                        <SidebarItem label="Logged In as {$currentSchool}">
                            <svelte:fragment slot="icon">
                                <ChartPieSolid
                                    class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                                />
                            </svelte:fragment>
                        </SidebarItem>
                        <SidebarItem label="Dashboard">
                            <svelte:fragment slot="icon">
                                <ChartPieSolid
                                    class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                                />
                            </svelte:fragment>
                        </SidebarItem>
                        <SidebarItem label="Add a Class" href="/dashboard/class"
                        ></SidebarItem>
                    {/if}
                </SidebarGroup>
            </div>

            <div class="mt-auto">
                <SidebarItem label="back to Webpage" href="/">
                    <svelte:fragment slot="icon">
                        <HomeSolid
                            class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                        />
                    </svelte:fragment>
                </SidebarItem>
            </div>
        </SidebarWrapper>
    </Sidebar>

    <div
        class="flex-grow transition-all duration-300 ease-in-out"
        class:ml-64={isOpen}
    >
        <slot></slot>
    </div>
</div>

<Button
    class="fixed top-4 left-4 z-20 transition-all duration-300 ease-in-out bg-blue-300 {isOpen
        ? 'hidden'
        : ''}"
    on:click={toggleSidebar}
>
    <BarsOutline class="w-5 h-5" />
</Button>
