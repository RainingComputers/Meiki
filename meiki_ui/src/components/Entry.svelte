<script context="module" lang="ts">
    function toCamelCase(spaceSeperatedString: string) {
        const words = spaceSeperatedString.toLowerCase().split(" ")

        const camelWords = words.map((word, index) => {
            if (index == 0) return word.toLowerCase()

            return word[0].toUpperCase() + word.slice(1)
        })

        return camelWords.join("")
    }
</script>

<script lang="ts">
    export let label: string
    export let password: boolean = false
    export let onEnter: () => void = undefined
    export let showLabel: boolean = true

    let inputEl: HTMLInputElement
    const type: string = password ? "password" : "text"

    const id: string = toCamelCase(label)

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

<div class=" flex flex-col w-full gap-1a">
    <label for={id} class="text-sm px-1" class:hidden={!showLabel}>
        {label}
    </label>

    <input
        bind:this={inputEl}
        {id}
        class=" bg-gray-200 focus:outline-none focus:border-blue-500 focus:ring-blue-500 focus:ring-2 rounded-lg p-2"
        {type}
        required
        on:keydown={onKeyDown}
    />
</div>
