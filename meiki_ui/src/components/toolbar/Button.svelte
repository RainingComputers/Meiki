<script lang="ts">
    import { createEventDispatcher } from "svelte"

    export let name: string = ""
    export let label: string = ""
    export let isButtonChecked: boolean = false

    const dispatch = createEventDispatcher()

    function onClick() {
        dispatch(name)
    }

    export function isChecked() {
        return isButtonChecked
    }
</script>

<div
    class="flex gap-2 flex-row items-center hover:bg-toolbar-focus rounded-xl p-2 cursor-pointer"
    class:isChecked={isButtonChecked}
    on:click={onClick}
    data-cy={name}
>
    {#if label}
        <span class="text-toolbar-content">{label}</span>
    {/if}

    <span class="stroke-toolbar-content h-6 w-6">
        <slot />
    </span>
</div>

<style>
    .isChecked {
        background: rgb(71 85 105);
    }
</style>
