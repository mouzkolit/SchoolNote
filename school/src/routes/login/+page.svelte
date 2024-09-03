<script lang="ts">
    import { onMount } from "svelte";

    interface School {
        id: number;
        name: string;
    }

    let selectedSchool: School | null = null;
    let schools: School[] = [];

    const getSchools = async () => {
        const response = await fetch("http://localhost:8080/schools");
        const data = await response.json();
        console.log(data);
        schools = data.schools.map((school: any) => ({
            id: school.ID,
            name: school.Name,
        }));
        console.log(schools);
    };

    onMount(async () => {
        await getSchools();
    });
</script>

<div class="flex justify-center items-center h-screen">
    <div class="bg-gray-100 rounded-lg p-6 shadow-md w-80">
        <h2 class="text-2xl font-bold text-center mb-6">School Selection</h2>
        <form class="flex flex-col">
            <label for="school" class="mb-2">Select School:</label>
            <select
                id="school"
                bind:value={selectedSchool}
                class="mb-4 p-2 border border-gray-300 rounded"
            >
                <option value={null} disabled>Choose a school</option>
                {#each schools as school (school.id)}
                    <option value={school}>{school.name}</option>
                {/each}
            </select>
        </form>

        {#if selectedSchool}
            <p class="text-center mt-4">You selected: {selectedSchool.name}</p>
        {/if}
        <button class="bg-blue-500 text-white p-2 rounded mt-4"
            >Login into Dashboard</button
        >
    </div>
</div>
