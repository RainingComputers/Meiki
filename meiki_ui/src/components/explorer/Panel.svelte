<script lang="ts">
    import Expander from "$cmp/Expander.svelte"
    import { fly } from "svelte/transition"

    export let onClick: () => void = undefined
    let panel: HTMLDivElement
</script>

<div class="flex flex-row h-full">
    <div
        class=" bg-base-1 flex flex-col h-full"
        style="min-width: 350px"
        data-cy="explorer"
        on:click={onClick}
        transition:fly|local={{ x: -400, duration: 200, opacity: 1 }}
        bind:this={panel}
    >
        <slot />
    </div>
    <Expander name="explorer-expander" attachedElement={panel} />
</div>
