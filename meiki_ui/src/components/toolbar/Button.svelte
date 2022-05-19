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
    class="flex gap-2 flex-row items-center hover:bg-slate-600 rounded-xl p-2"
    class:isChecked={isButtonChecked}
    on:click={onClick}
    data-cy={name}
>
    {#if label}
        <span class="text-gray-200">{label}</span>
    {/if}

    <span class="stroke-gray-200 h-6 w-6">
        <slot />
    </span>
</div>

<style>
    .isChecked {
        background: rgb(71 85 105);
    }
</style>
