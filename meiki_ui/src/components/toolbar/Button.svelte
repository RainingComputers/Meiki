<script lang="ts">
    import { createEventDispatcher } from "svelte"

    export let checkable: boolean = false
    export let name: string = ""
    export let showLabel: boolean = false
    let isButtonChecked: boolean = true

    const dispatch = createEventDispatcher()

    function onClick() {
        if (checkable) isButtonChecked = !isButtonChecked
        dispatch(name)
    }

    export function isChecked() {
        return isButtonChecked
    }
</script>

<div
    class="flex gap-2 flex-row items-center hover:bg-slate-600 rounded-xl p-2"
    class:isChecked={isButtonChecked && checkable}
    on:click={onClick}
    data-cy={name}
>
    {#if showLabel}
        <span class="text-gray-200">{name}</span>
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
