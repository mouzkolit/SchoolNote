<script>
    import "../app.css";
    import { Navbar, NavLi, NavUl } from "flowbite-svelte";
</script>

<div class="relative px-8">
    <Navbar
        class="bg-gray-200 text-white px-4 sm:px-6 py-3 fixed w-full z-20 top-0 left-0 border-b border-gray-100 shadow-lg"
    >
        <div class="flex items-center justify-between">
            <NavUl class="flex space-x-8 ml-auto">
                <NavLi
                    href="/"
                    class="hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium"
                >
                    Home
                </NavLi>
                <NavLi
                    href="/school"
                    active={true}
                    class="hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium"
                >
                    Add School
                </NavLi>
                <NavLi
                    href="/login"
                    class="hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium"
                >
                    Login
                </NavLi>
                <NavLi
                    href="/pricing"
                    class="hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium"
                >
                    Pricing
                </NavLi>
                <NavLi
                    href="/contact"
                    class="hover:text-gray-300 px-3 py-2 rounded-md text-sm font-medium"
                    >Contact</NavLi
                >
            </NavUl>
        </div>
    </Navbar>
</div>
<div class="mt-20">
    <slot />
</div>
