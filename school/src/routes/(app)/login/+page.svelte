<script lang="ts">
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { schoolStore } from "$lib/schoolStore";

    interface School {
        ID: number;
        Name: string;
    }

    let selectedSchool: School | null = null;
    let schools: School[] = [];

    const getSchools = async () => {
        const response = await fetch("http://localhost:8080/schools", {
            credentials: "include",
        });
        console.log(response);
        const data = await response.json();
        console.log(data);
        schools = data.schools.map((school: any) => ({
            ID: school.ID,
            Name: school.Name,
        }));
    };

    const goToDashboard = () => {
        if (selectedSchool) {
            console.log(selectedSchool);
            goto("/dashboard/", {
                state: selectedSchool,
            });
        } else {
            alert("Please select a school before proceeding to the dashboard.");
        }
    };

    $: if (selectedSchool) {
        schoolStore.set({
            schoolName: selectedSchool.Name,
            schoolId: selectedSchool.ID.toString(),
        });
    }

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
                {#each schools as school (school.ID)}
                    <option value={school}>{school.Name}</option>
                {/each}
            </select>
        </form>

        {#if selectedSchool}
            <p class="text-center mt-4">You selected: {selectedSchool.Name}</p>
        {/if}
        <button
            class="bg-blue-500 text-white p-2 rounded mt-4 w-full hover:bg-blue-600"
            on:click={goToDashboard}
        >
            Go to Dashboard
        </button>
    </div>
</div>
