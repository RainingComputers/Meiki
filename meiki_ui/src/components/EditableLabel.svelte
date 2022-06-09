<script lang="ts">
    // TODO redesign this component

    export let text: string = ""
    export let onEnter: (text: string) => void

    let editableMode: boolean = false

    let spanElement: HTMLSpanElement

    function onKeyDown(event: any) {
        const newText = spanElement.innerText.trim()
        // TODO: Show error if empty
        // TODO: Sometimes click doesn't register first time
        if (!onEnter) return
        if (event.type !== "focusout" && event.key !== "Enter") return

        editableMode = false
        if (text !== newText) {
            onEnter(newText)
        }
    }
</script>

<span
    contenteditable={editableMode ? "true" : "false"}
    bind:this={spanElement}
    on:keydown={onKeyDown}
    on:focusout={onKeyDown}
    on:click={() => {
        editableMode = true
        spanElement.focus()
    }}
>
    {text}
</span>
