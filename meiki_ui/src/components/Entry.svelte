<script lang="ts">
    export let label: string
    export let password: boolean = false
    export let onEnter: () => void = undefined

    let inputEl: HTMLInputElement
    const type: string = password ? "password" : "text"

    const id: string = label.toLowerCase().replace(" ", "")

    function onKeyDown(event: any) {
        if (!onEnter) return
        if (event.key !== "Enter" || event.target.value.length === 0) return
        onEnter()
    }

    export function getValue() {
        return inputEl.value
    }

    export function focus() {
        inputEl.focus()
    }
</script>

<div class=" flex flex-col w-full gap-1">
    <label for={id} class="text-sm px-1">{label}</label>

    <input
        bind:this={inputEl}
        {id}
        class=" bg-gray-200 focus:outline-none focus:border-blue-500 focus:ring-blue-500 focus:ring-2 rounded-lg p-2"
        {type}
        required
        on:keydown={onKeyDown}
    />
</div>
