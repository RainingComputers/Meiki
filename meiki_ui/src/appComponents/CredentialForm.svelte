<script lang="ts">
    import Card from "../components/Card.svelte"
    import Button from "../components/Button.svelte"
    import Entry from "../components/Entry.svelte"
    import Logo from "../components/Logo.svelte"

    export let type: string = "login"
</script>

<div class="flex flex-row">
    <Card>
        <Logo />
        <div class="flex flex-col w-full gap-2">
            <Entry label="Username" />
            <Entry label="Password" />
        </div>
        {#if type === "login"}
            <Button expand={true} label="Login" />
        {:else}
            <Button expand={true} label="Create Meiki account" />
        {/if}
    </Card>
</div>
