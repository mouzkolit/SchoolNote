<script>
    import { Label, Input } from "flowbite-svelte";

    let schoolName = "";
    let schoolPlace = "";
    let schoolType = "";
    let schoolWebsite = "";

    const handleSubmit = () => {
        submitSchool();
    };

    async function submitSchool() {
        var url = new URL("http://localhost:8080/school");
        var params = {
            name: schoolName,
            place: schoolPlace,
            web: schoolWebsite,
        };
        url.search = new URLSearchParams(params).toString();
        const response = await fetch(url, {
            method: "POST",
        });
    }
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
                    bind:value={schoolName}
                />
            </div>
            <div class="mb-6">
                <Label for="school-place" class="block mb-2">School Place</Label
                >
                <Input
                    id="school-place"
                    placeholder="Default input"
                    class="border p-2"
                    bind:value={schoolPlace}
                />
            </div>
            <div class="mb-6">
                <Label for="school-type" class="block mb-2">School Type</Label>
                <Input
                    id="school-type"
                    placeholder="Default input"
                    class="border p-2"
                    bind:value={schoolType}
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
                    bind:value={schoolWebsite}
                />
            </div>
            <button
                class="mt-4 bg-blue-500 text-white px-4 py-2 rounded"
                on:click={handleSubmit}
            >
                Add School
            </button>
        </div>
    </div>
</div>
