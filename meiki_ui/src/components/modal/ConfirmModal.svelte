<script lang="ts">
    import ActionModal from "$cmp/modal/ActionModal.svelte"
    import Button, { ButtonType } from "$cmp/Button.svelte"

    export let error: string = ""
    export let message: string
    export let onYes: () => void
    export let onNo: () => void
</script>

<ActionModal {error}>
    <div class="flex flex-col gap-3">
        <p class="px-1">{message}</p>
        <div class="flex gap-2 w-full">
            <Button expand={true} label="Yes" onClick={onYes} />
            <Button
                expand={true}
                label="No"
                type={ButtonType.SECONDARY}
                onClick={onNo}
            />
        </div>
    </div>
</ActionModal>
