<script>
    import { page } from "$app/stores";
    import { isAuthenticated, currentSchool } from "$lib/authStore";
    import { schoolStore } from "$lib/schoolStore";

    import {
        Sidebar,
        SidebarGroup,
        SidebarItem,
        SidebarWrapper,
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

    let activeClass =
        "flex items-center p-2 text-base font-normal text-primary-900 bg-primary-200 dark:bg-primary-700 rounded-lg dark:text-white hover:bg-primary-100 dark:hover:bg-gray-700";
    let nonActiveClass =
        "flex items-center p-2 text-base font-normal text-white-100 rounded-lg dark:text-white hover:bg-green-100 dark:hover:bg-green-700";

    export let isOpen = true;
    const toggleSidebar = () => (isOpen = !isOpen);
</script>

<div class="flex h-screen w-full relative">
    <div
        class="transition-all duration-300 ease-in-out {isOpen
            ? 'w-64'
            : 'w-16'}"
    >
        <Sidebar {activeUrl} class="h-full" {activeClass} {nonActiveClass}>
            <SidebarWrapper class="bg-gray-800 text-white h-full flex flex-col">
                <div
                    class="flex items-center justify-between p-4 border-b border-gray-700"
                >
                    {#if isOpen}
                        <span class="text-2xl font-bold">SchoolNote</span>
                    {/if}
                    <Button
                        class="p-1 hover:bg-gray-700 rounded-full"
                        color="none"
                        on:click={toggleSidebar}
                    >
                        <BarsOutline class="w-6 h-6" />
                    </Button>
                </div>
                {#if isOpen}
                    <div class="p-4 border-b border-gray-700">
                        <p class="text-sm text-gray-400">School</p>
                        <p class="font-semibold">{schoolName}</p>
                    </div>
                {/if}
                <SidebarGroup class="flex-grow">
                    {#if !$isAuthenticated}
                        <SidebarItem
                            label={isOpen ? "Login" : ""}
                            href="/dashboard/login"
                            color="white"
                        >
                            <svelte:fragment slot="icon">
                                <ArrowRightToBracketOutline class="w-5 h-5" />
                            </svelte:fragment>
                        </SidebarItem>
                    {:else}
                        <SidebarItem
                            label={isOpen ? "Dashboard" : ""}
                            href="/dashboard"
                            color="white"
                        >
                            <svelte:fragment slot="icon">
                                <ChartPieSolid class="w-5 h-5" />
                            </svelte:fragment>
                        </SidebarItem>
                        <SidebarItem
                            label={isOpen ? "Classes" : ""}
                            href="/dashboard/class"
                            color="white"
                        >
                            <svelte:fragment slot="icon">
                                <HomeSolid class="w-5 h-5" />
                            </svelte:fragment>
                        </SidebarItem>
                    {/if}
                </SidebarGroup>
                <div class="p-4 border-t border-gray-700">
                    <SidebarItem
                        label={isOpen ? "Home" : ""}
                        href="/"
                        color="white"
                    >
                        <svelte:fragment slot="icon">
                            <HomeSolid class="w-5 h-5" />
                        </svelte:fragment>
                    </SidebarItem>
                </div>
            </SidebarWrapper>
        </Sidebar>
    </div>

    <div class="flex-grow flex flex-col h-full bg-gray-100 overflow-x-hidden">
        <div class="p-4 flex-grow">
            <div class="mt-12">
                <slot />
            </div>
        </div>
    </div>
</div>
