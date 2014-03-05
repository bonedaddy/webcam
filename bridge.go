package main

// #cgo CXXFLAGS: -DV8_DEPRECATION_WARNINGS -DEXPAT_RELATIVE_PATH -DFEATURE_ENABLE_VOICEMAIL -DGTEST_RELATIVE_PATH -DJSONCPP_RELATIVE_PATH -DLOGGING=1 -DSRTP_RELATIVE_PATH -DFEATURE_ENABLE_SSL -DFEATURE_ENABLE_PSTN -DHAVE_SRTP -DHAVE_WEBRTC_VIDEO -DHAVE_WEBRTC_VOICE -DUSE_WEBRTC_DEV_BRANCH -D_FILE_OFFSET_BITS=64 -DCHROMIUM_BUILD -DTOOLKIT_VIEWS=1 -DUI_COMPOSITOR_IMAGE_TRANSPORT -DUSE_AURA=1 -DUSE_CAIRO=1 -DUSE_GLIB=1 -DUSE_DEFAULT_RENDER_THEME=1 -DUSE_LIBJPEG_TURBO=1 -DUSE_NSS=1 -DUSE_X11=1 -DUSE_CLIPBOARD_AURAX11=1 -DENABLE_ONE_CLICK_SIGNIN -DUSE_XI2_MT=2 -DENABLE_REMOTING=1 -DENABLE_WEBRTC=1 -DENABLE_PEPPER_CDMS -DENABLE_CONFIGURATION_POLICY -DENABLE_INPUT_SPEECH -DENABLE_NOTIFICATIONS -DUSE_UDEV -DENABLE_EGLIMAGE=1 -DENABLE_TASK_MANAGER=1 -DENABLE_EXTENSIONS=1 -DENABLE_PLUGIN_INSTALLATION=1 -DENABLE_PLUGINS=1 -DENABLE_SESSION_SERVICE=1 -DENABLE_THEMES=1 -DENABLE_AUTOFILL_DIALOG=1 -DENABLE_BACKGROUND=1 -DENABLE_AUTOMATION=1 -DENABLE_GOOGLE_NOW=1 -DCLD_VERSION=2 -DENABLE_FULL_PRINTING=1 -DENABLE_PRINTING=1 -DENABLE_SPELLCHECK=1 -DENABLE_CAPTIVE_PORTAL_DETECTION=1 -DENABLE_APP_LIST=1 -DENABLE_SETTINGS_APP=1 -DENABLE_MANAGED_USERS=1 -DENABLE_MDNS=1 -DLIBPEERCONNECTION_LIB=1 -DLINUX -DHAVE_SCTP -DHASH_NAMESPACE=__gnu_cxx -DPOSIX -DDISABLE_DYNAMIC_CAST -D_REENTRANT -DSSL_USE_NSS -DHAVE_NSS_SSL_H -DSSL_USE_NSS_RNG -DNDEBUG -DNVALGRIND -DDYNAMIC_ANNOTATIONS_ENABLED=0
// #cgo CXXFLAGS: -pthread
// #cgo CXXFLAGS: -I/home/s/src/webrtc.org/trunk  -I/home/s/src/webrtc.org/trunk/third_party  -I/home/s/src/webrtc.org/trunk/third_party/webrtc  -I/home/s/src/webrtc.org/trunk/webrtc -I/home/s/src/webrtc.org/trunk/net/third_party/nss/ssl  -I/home/s/src/webrtc.org/trunk/third_party/jsoncpp/overrides/include  -I/home/s/src/webrtc.org/trunk/third_party/jsoncpp/source/include
// 
// #cgo LDFLAGS: -lstdc++ -lm -lnss3 -lnssutil3 -lX11 -lXext -lcrypto -lplc4 -lnspr4 -lexpat -ldl
// #cgo LDFLAGS: -Wl,--start-group /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/jsoncpp/libjsoncpp.a /home/s/src/webrtc.org/trunk/out/Release/obj/talk/libjingle_peerconnection.a /home/s/src/webrtc.org/trunk/out/Release/obj/net/third_party/nss/libcrssl.a /home/s/src/webrtc.org/trunk/out/Release/obj/talk/libjingle.a /home/s/src/webrtc.org/trunk/out/Release/obj/talk/libjingle_media.a /home/s/src/webrtc.org/trunk/out/Release/libyuv.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libjpeg_turbo/libjpeg_turbo.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/usrsctp/libusrsctplib.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libvideo_capture_module.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libwebrtc_utility.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libaudio_coding_module.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libCNG.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/common_audio/libcommon_audio.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/system_wrappers/source/libsystem_wrappers.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/common_audio/libcommon_audio_sse2.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libG711.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libG722.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libiLBC.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libiSAC.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libiSACFix.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libPCM16B.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libNetEq.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libwebrtc_opus.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/opus/libopus.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libacm2.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libNetEq4.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libmedia_file.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libwebrtc_video_coding.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libwebrtc_i420.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/common_video/libcommon_video.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/video_coding/utility/libvideo_coding_utility.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/video_coding/codecs/vp8/libwebrtc_vp8.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libvpx/libvpx.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libvpx/libvpx_asm_offsets_vp8.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libvpx/libvpx_intrinsics_mmx.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libvpx/libvpx_intrinsics_sse2.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libvpx/libvpx_intrinsics_ssse3.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libvideo_render_module.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/video_engine/libvideo_engine_core.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/librtp_rtcp.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libpaced_sender.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libremote_bitrate_estimator.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/remote_bitrate_estimator/librbe_components.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libbitrate_controller.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libvideo_processing.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libvideo_processing_sse2.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/voice_engine/libvoice_engine.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libaudio_conference_mixer.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libaudio_processing.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libaudioproc_debug_proto.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/protobuf/libprotobuf_lite.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libaudio_processing_sse2.a /home/s/src/webrtc.org/trunk/out/Release/obj/webrtc/modules/libaudio_device.a /home/s/src/webrtc.org/trunk/out/Release/obj/talk/libjingle_sound.a /home/s/src/webrtc.org/trunk/out/Release/obj/talk/libjingle_p2p.a /home/s/src/webrtc.org/trunk/out/Release/obj/third_party/libsrtp/libsrtp.a -Wl,--end-group
import "C"