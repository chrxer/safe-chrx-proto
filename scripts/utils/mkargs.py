#!/usr/bin/python3

from typing import Dict

# https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/etc/apikeys
API_KEYS = {
    "google_api_key":"AIzaSyCkfPOPZXDKNn8hhgu3JrA62wIgC93d44k",
    "google_default_client_id":"811574891467.apps.googleusercontent.com",
    "google_default_client_secret":"kdloedMFGdGla2P1zacGjAQh",
}


def make(debug=False, gapi_keys:Dict[str]=API_KEYS):

    if debug:
        dstr = "true"
        offbuild="false"
        symb=2
    else:
        dstr = "false"
        offbuild="true"
        symb=0

    args = [
        # https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/rules
        "clang_use_chrome_plugins=false",

        # disabled features
        "use_libjpeg_turbo=true"
        "use_unofficial_version_number=false",
        "safe_browsing_use_unrar=false",
        "enable_vr=false",
        "enable_nacl=false",
        "enable_reading_list=false",
        "enable_iterator_debugging=false",
        "enable_hangout_services_extension=false",
        "angle_has_histograms=false",
        "build_angle_perftests=false",
        "treat_warnings_as_errors=false",
        "use_qt=false",
        "is_cfi=false",
        "chrome_pgo_phase=0",

        # enabled features
        "use_gio=true",
        f"is_official_build={offbuild}",
        "use_pulseaudio=true",
        "link_pulseaudio=true",
        "rtc_use_pipewire=true",
        "icu_use_data_file=true",
        "enable_widevine=true",
        "v8_enable_backtrace=true",
        "proprietary_codecs=true",
        "ffmpeg_branding=\"Chrome\"",
        "disable_fieldtrial_testing_config=true",

        "cc_wrapper=\"env CCACHE_SLOPPINESS=time_macros CCACHE_NOHASHDIR=1 CCACHE_LOGFILE=/tmp/ccache_log.log ccache\""
    ]

    # set symbol levels
    for arg in ["symbol_level","blink_symbol_level","v8_symbol_level"]:
        args.append(f"{arg}={symb}")
    
    # enable//disable tests//debugging
    for arg in ["is_debug", "build_dawn_tests", "angle_build_tests"]:
        args.append(f"{arg}={dstr}")

    # google API keys
    for key, value in gapi_keys.items():
        args.append(f"{key}=\"{dstr}\"")