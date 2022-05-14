<script lang="ts">
    import { onMount } from "svelte"
    import Entry from "$cmp/Entry.svelte"
    import ActionModal from "$cmp/modal/ActionModal.svelte"
    import Button from "$cmp/Button.svelte"

    export let error: string = ""
    export let entryLabel: string
    export let buttonLabel: string
    export let onValue: (value: string) => any

    let entry: Entry

    onMount(() => {
        entry.focus()
    })

    function onClickOrOnEnter() {
        onValue(entry.getValue())
    }
</script>

<ActionModal {error}>
    <Entry label={entryLabel} onEnter={onClickOrOnEnter} bind:this={entry} />
    <Button label={buttonLabel} onClick={onClickOrOnEnter} fullWidth={true} />
</ActionModal>
