<script lang="ts">
    import Root from "../components/Root.svelte"
    import Card from "../components/Card.svelte"
    import Button from "../components/Button.svelte"
    import Entry from "../components/Entry.svelte"
    import Logo from "../components/Logo.svelte"
    import Link from "../components/Link.svelte"
</script>

<Root>
    <div class="flex flex-row">
        <Card>
            <Logo />
            <div class="flex flex-col w-full gap-2">
                <Entry label="Username" />
                <Entry label="Password" />
            </div>
            <Button expand={true} label="Login" />
        </Card>
    </div>
    <div class=" mt-8 text-sm">
        New to Meiki notes? <Link text="Create" /> a new account to get started
    </div>
</Root>
