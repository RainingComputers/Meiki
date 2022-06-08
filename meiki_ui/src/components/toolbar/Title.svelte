<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import EditableLabel from "$cmp/EditableLabel.svelte"

    const dispatchEvent = createEventDispatcher()

    export let title: string

    function onEnter(text: string) {
        dispatchEvent("rename", {newTitle: text})
    }
</script>

<span
    class="flex flex-row flex-grow text-gray-200 justify-center font-light items-center"
>
    <EditableLabel text={title} {onEnter} />
    <slot />
</span>
