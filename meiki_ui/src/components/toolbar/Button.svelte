<script lang="ts">
    export let checked: boolean = false
    export let label: string = ""
    let isButtonChecked: boolean = false

    function onClick() {
        if (checked) isButtonChecked = !isButtonChecked
    }

    export function isChecked() {
        return isButtonChecked
    }
</script>

<div
    class="flex gap-2 flex-row items-center hover:bg-slate-600 rounded-xl p-2"
    class:bg-slate-600={isButtonChecked}
    on:click={onClick}
>
    {#if label}
        <span class="text-gray-200">{label}</span>
    {/if}

    <span class="stroke-gray-200 h-6 w-6">
        <slot />
    </span>
</div>
