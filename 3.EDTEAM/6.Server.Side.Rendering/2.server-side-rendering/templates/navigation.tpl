{{ define "navigation" }}
<header class="main-header">
    <div class="ed-grid s-grid-5 lg-grid-4">
        <div class="s-cols-4 lg-cols-1 s-cross-center">
            <a href="/">
                <img src="img/logo.svg" alt="" class="main-logo">
            </a>
        </div>
        <div class="s-grid-1 lg-cols-3 s-cross-center s-main-end header-links">
            <nav class="main-menu" id="main-menu">
                <ul>
                    <li><a href="/" target="_blank">Home</a></li>
                    <li><a href="https://ed.team">Go to ED Team</a></li>
                </ul>
            </nav>
            <div class="main-menu-toggle to-l" id="main-menu-toggle"></div>
        </div>
    </div>
</header>
{{ end }}