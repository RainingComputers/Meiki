<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import Searchbar from "$cmp/Searchbar.svelte"
    import Button from "$cmp/Button.svelte"
    import Toolbar from "$cmp/toolbar/Toolbar.svelte"

    const dispatchEvent = createEventDispatcher()
    let width = "25%"

    export function setWidth(newWidth: string) {
        width = newWidth
    }
</script>

<Toolbar {width} expand={false}>
    <Searchbar on:change />
    <Button label="New" onClick={() => dispatchEvent("createNote")} />
</Toolbar>
