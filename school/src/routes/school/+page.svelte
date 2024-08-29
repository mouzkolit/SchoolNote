<script>
    import { Label, Input } from "flowbite-svelte";
</script>

<div class="flex items-center justify-center min-h-screen">
    <div class="bg-gray-200 p-8 rounded shadow-lg">
        <div>
            <div class="mb-6">
                <Label for="school-name" class="block mb-2">School Name</Label>
                <Input
                    id="school-name"
                    placeholder="Default input"
                    class="border p-2"
                />
            </div>
            <div class="mb-6">
                <Label for="school-place" class="block mb-2">School Place</Label
                >
                <Input
                    id="school-place"
                    placeholder="Default input"
                    class="border p-2"
                />
            </div>
            <div class="mb-6">
                <Label for="school-type" class="block mb-2">School Type</Label>
                <Input
                    id="school-type"
                    placeholder="Default input"
                    class="border p-2"
                />
            </div>
            <div class="mb-6">
                <Label for="school-website" class="block mb-2"
                    >School Website</Label
                >
                <Input
                    id="school-website"
                    placeholder="Default input"
                    class="border p-2"
                />
            </div>
            <button class="mt-4 bg-blue-500 text-white px-4 py-2 rounded">
                Add School
            </button>
        </div>
    </div>
</div>
