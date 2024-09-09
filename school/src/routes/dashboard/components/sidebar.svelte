<script>
    import { page } from "$app/stores";
    import { isAuthenticated, currentSchool } from "$lib/authStore";
    import {
        Sidebar,
        SidebarGroup,
        SidebarItem,
        SidebarWrapper,
        Heading,
        Span,
    } from "flowbite-svelte";
    import {
        ChartPieSolid,
        ArrowRightToBracketOutline,
        HomeSolid,
    } from "flowbite-svelte-icons";
    $: activeUrl = $page.url.pathname;
</script>

<Sidebar {activeUrl}>
    <SidebarWrapper class="bg-blue-200 h-screen fixed left-0 top-0 flex-col">
        <div>
            <Heading
                tag="h1"
                class="mb-4 p-4"
                customSize="text-3xl font-extrabold md:text-5xl lg:text-6xl"
            >
                <Span gradient>Admin</Span>.
            </Heading>
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
                    <SidebarItem label="Add a Class"></SidebarItem>
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
