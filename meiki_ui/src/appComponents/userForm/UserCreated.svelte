<script>
    import Link from "../../components/Link.svelte"
    import Root from "../../components/Root.svelte"
    import Success from "../../components/Success.svelte"
</script>

<Root>
    <div class="flex flex-col items-center gap-5 ">
        <Success>Your account has successfully been created</Success>
        <span>
            Click <Link target="/login" text="here" /> to login
        </span>
    </div>
</Root>
