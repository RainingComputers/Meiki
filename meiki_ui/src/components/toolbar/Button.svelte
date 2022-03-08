<script lang="ts">
    export let isCheckedButton: boolean = false
    export let label: string = ""
    let checked: boolean = false

    function onClick() {
        if (isCheckedButton) checked = !checked
    }

    export function isChecked() {
        return checked
    }
</script>

<div
    class="flex flex-row items-center hover:bg-slate-700 rounded-xl p-2"
    class:bg-slate-700={checked}
>
    {#if label}
        <span class="text-gray-200">{label}</span>
    {/if}

    <span on:click={onClick} class="stroke-gray-200 h-6 w-6">
        <slot />
    </span>
</div>
