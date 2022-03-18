<script lang="ts">
    import Root from "$cmp/Root.svelte"
    import Link from "$cmp/Link.svelte"
</script>

<Root>
    <div
        class=" grid h-screen w-full v-screen font-serif text-gray-800 bg-gray-200 place-content-center dark:bg-gray-800 dark:text-gray-50 sm:p-0"
    >
        <h1 class="text-9xl">Meiki UI</h1>
        <div class="mt-8 text-sm">
            New to Meiki notes? <Link text="Create" target="/create" /> a new account
            to get started
        </div>
        <div class="mt-8 text-sm">
            If you already have an account, Click <Link
                text="here"
                target="/login"
            /> to login
        </div>
    </div>
</Root>
