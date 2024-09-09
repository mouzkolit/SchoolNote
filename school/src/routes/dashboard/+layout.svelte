<script>
    import "../../app.css";
    import Sidebar from "./components/sidebar.svelte";

    let isOpen = false;
</script>

<Sidebar bind:isOpen>
    <main class="p-4">
        <!-- Your main dashboard content goes here -->
        <slot></slot>
    </main>
</Sidebar>
