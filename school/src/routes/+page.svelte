<script>
    import {
        Section,
        News,
        HeroHeader,
        HeroBody,
    } from "flowbite-svelte-blocks";
    import { Button } from "flowbite-svelte";
    import { ArrowRightOutline } from "flowbite-svelte-icons";
</script>

<Section name="heroVisual">
    <div class="mr-auto place-self-center lg:col-span-7">
        <HeroHeader
            h1Class="max-w-2xl mb-4 text-4xl font-extrabold tracking-tight leading-none md:text-5xl xl:text-6xl dark:text-white"
            pClass="max-w-2xl mb-6 font-light text-gray-500 lg:mb-8 md:text-lg lg:text-xl dark:text-gray-400"
        >
            <svelte:fragment slot="h1"
                >School Software for Parents, Pupils and Teachers</svelte:fragment
            >
            <svelte:fragment slot="paragraph"
                >Be more connected, education is everything!</svelte:fragment
            >
            <a href="/"
                ><Button
                    size="xl"
                    color="red"
                    class="inline-flex items-center justify-center mr-3"
                    >Connect your School<ArrowRightOutline
                        size="md"
                        class="ml-2 -mr-1"
                    /></Button
                >
            </a>
        </HeroHeader>
    </div>
</Section>
