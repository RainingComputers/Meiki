<script lang="ts">

    // TODO redesign this component

    export let text: string = ""
    export let onEnter: (text: string) => void

    let editableMode: boolean = false

    let spanElement: HTMLSpanElement

    function onKeyDown(event: any) {
        // TODO: move to utils
        if (!onEnter) return
        if (event.key !== "Enter") return
        editableMode = false
        onEnter(event.target.innerText.trim())
    }

    function onFocusout() {
        // spanElement.innerText = text
        editableMode = false
        spanElement.innerText = text
    }
</script>

<span
    contenteditable={editableMode ? "true" : "false"}
    class=" inline-block"
    bind:this={spanElement}
    on:keydown={onKeyDown}
    on:focusout={onFocusout}
    on:click={() => {
        editableMode = true
        spanElement.focus()
    }}
>
    {text}
</span>
