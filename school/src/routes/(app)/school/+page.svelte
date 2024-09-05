<script>
    import { Label, Input } from "flowbite-svelte";
    import schoolLogo from "$lib/assets/schoolnote.jpg";

    let schoolName = "";
    let schoolPlace = "";
    let schoolType = "";
    let schoolWebsite = "";
    let password = "";

    const handleSubmit = () => {
        submitSchool();
    };

    async function submitSchool() {
        var url = new URL("http://localhost:8080/school");
        var params = {
            name: schoolName,
            place: schoolPlace,
            web: schoolWebsite,
            password: password,
        };
        url.search = new URLSearchParams(params).toString();
        const response = await fetch(url, {
            method: "POST",
        });
        const data = await response.json();
        console.log(data);
    }
</script>

<div class="relative min-h-screen flex items-center justify-center">
    <div class="absolute inset-0 z-0">
        <img
            src={schoolLogo}
            alt="School Background"
            class="w-full h-full object-cover"
        />
        <div class="absolute inset-0 bg-black opacity-50"></div>
    </div>

    <div
        class="relative z-10 bg-white bg-opacity-90 p-8 rounded-lg shadow-xl max-w-md w-full"
    >
        <h2 class="text-3xl font-bold text-center mb-6">Add New School</h2>
        <form>
            <div class="mb-4">
                <Label for="school-name" class="block mb-2">School Name</Label>
                <Input
                    id="school-name"
                    placeholder="Enter school name"
                    class="border p-2 w-full"
                    bind:value={schoolName}
                />
            </div>
            <div class="mb-4">
                <Label for="password" class="block mb-2">Password</Label>
                <Input
                    id="password"
                    placeholder="Enter your password"
                    class="border p-2 w-full"
                    bind:value={password}
                />
            </div>
            <div class="mb-4">
                <Label for="school-place" class="block mb-2">School Place</Label
                >
                <Input
                    id="school-place"
                    placeholder="Enter school location"
                    class="border p-2 w-full"
                    bind:value={schoolPlace}
                />
            </div>
            <div class="mb-4">
                <Label for="school-type" class="block mb-2">School Type</Label>
                <Input
                    id="school-type"
                    placeholder="Enter school type"
                    class="border p-2 w-full"
                    bind:value={schoolType}
                />
            </div>
            <div class="mb-6">
                <Label for="school-website" class="block mb-2"
                    >School Website</Label
                >
                <Input
                    id="school-website"
                    placeholder="Enter school website"
                    class="border p-2 w-full"
                    bind:value={schoolWebsite}
                />
            </div>
            <button
                class="w-full bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition duration-300"
                on:click={handleSubmit}
            >
                Add School
            </button>
        </form>
    </div>
</div>
