<script lang="ts">
    import Root from "../components/Root.svelte"
    import Card from "../components/Card.svelte"
    import Button from "../components/Button.svelte"
</script>

<Root>
    <div class="flex flex-row">
        <Card>
            Meiki UI
            <Button expand={true} label="Login" />
        </Card>
        
    </div>
</Root>
