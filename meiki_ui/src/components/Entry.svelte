<script lang="ts">
    export let label: string = "label"
    export let password: boolean = false

    let inputEl: HTMLInputElement
    const type: string = password ? "password" : "text"

    export function getValue() {
        return inputEl.value
    }
</script>

<div class=" flex flex-col w-full gap-1">
    <span class=" text-sm">{label}</span>

    <input
        bind:this={inputEl}
        class=" bg-gray-300 rounded-lg p-2"
        {type}
        required
    />
</div>
