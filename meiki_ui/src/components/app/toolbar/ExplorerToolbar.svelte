<script lang="ts">
    import Entry from "$cmp/Entry.svelte"
    import { createEventDispatcher } from "svelte"

    const dispatchEvent = createEventDispatcher()
</script>

<div class="flex flex-row py-2 px-4 gap-4">
    <Entry showLabel={false} label="search" />
    <div
        class="  hover:cursor-pointer m-auto "
        on:click={() => dispatchEvent("create")}
    >
        +
    </div>
</div>
