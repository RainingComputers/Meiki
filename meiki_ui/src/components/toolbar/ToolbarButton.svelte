<script lang="ts">
    export let isCheckedButton: boolean = false

    let checked: boolean = false

    function onClick() {
        if (isCheckedButton) checked = !checked
    }

    export function isChecked() {
        return checked
    }
</script>

<span
    on:click={onClick}
    class:bg-slate-700={checked}
    class="stroke-gray-200 h-9 w-9 hover:bg-slate-700 p-2 rounded-xl"
>
    <slot />
</span>
