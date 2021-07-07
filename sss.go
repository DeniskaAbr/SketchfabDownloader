package main

// #js-dom-data-prefetched-data
// ss.childNodes[0].data

type MyJsonName struct {
	I_categories struct {
		Count   int64 `json:"count"`
		Results []struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
			UID  string `json:"uid"`
		} `json:"results"`
	} `json:"/i/categories"`
	I_models_29c76ea294264212b0598f358fbce111 struct {
		AnimationCount int64  `json:"animationCount"`
		Description    string `json:"description"`
		DisplayStatus  string `json:"displayStatus"`
		DownloadType   string `json:"downloadType"`
		EditorURL      string `json:"editorUrl"`
		EmbedURL       string `json:"embedUrl"`
		FaceCount      int64  `json:"faceCount"`
		Files          []struct {
			Flag          int64  `json:"flag"`
			ModelSize     int64  `json:"modelSize"`
			OsgjsSize     int64  `json:"osgjsSize"`
			OsgjsURL      string `json:"osgjsUrl"`
			UID           string `json:"uid"`
			WireframeSize int64  `json:"wireframeSize"`
		} `json:"files"`
		HasEnterpriseTracking bool        `json:"hasEnterpriseTracking"`
		InStore               bool        `json:"inStore"`
		IsArEnabled           bool        `json:"isArEnabled"`
		IsDeleted             bool        `json:"isDeleted"`
		IsDisabled            bool        `json:"isDisabled"`
		IsDownloadable        bool        `json:"isDownloadable"`
		IsInspectable         bool        `json:"isInspectable"`
		IsPublished           bool        `json:"isPublished"`
		IsRestricted          bool        `json:"isRestricted"`
		License               interface{} `json:"license"`
		LikeCount             int64       `json:"likeCount"`
		Name                  string      `json:"name"`
		Options               struct {
			Animation struct {
				CycleMode              string        `json:"cycleMode"`
				InitializeWithRestPose bool          `json:"initializeWithRestPose"`
				Order                  []interface{} `json:"order"`
				Speed                  int64         `json:"speed"`
			} `json:"animation"`
			Background struct {
				Color  []int64 `json:"color"`
				Enable string  `json:"enable"`
				UID    string  `json:"uid"`
			} `json:"background"`
			Camera struct {
				Down                 float64   `json:"down"`
				Fov                  int64     `json:"fov"`
				Left                 float64   `json:"left"`
				NearFarRatio         float64   `json:"nearFarRatio"`
				Position             []float64 `json:"position"`
				Right                float64   `json:"right"`
				Target               []float64 `json:"target"`
				Up                   float64   `json:"up"`
				UseCameraConstraints bool      `json:"useCameraConstraints"`
				UsePanConstraints    bool      `json:"usePanConstraints"`
				UsePitchConstraints  bool      `json:"usePitchConstraints"`
				UseYawConstraints    bool      `json:"useYawConstraints"`
				UseZoomConstraints   bool      `json:"useZoomConstraints"`
				ZoomIn               int64     `json:"zoomIn"`
				ZoomOut              int64     `json:"zoomOut"`
			} `json:"camera"`
			CreatedAt   string `json:"createdAt"`
			Environment struct {
				BackgroundExposure int64   `json:"backgroundExposure"`
				Blur               float64 `json:"blur"`
				Enable             bool    `json:"enable"`
				Exposure           int64   `json:"exposure"`
				LightIntensity     int64   `json:"lightIntensity"`
				Rotation           int64   `json:"rotation"`
				ShadowBias         float64 `json:"shadowBias"`
				ShadowEnable       bool    `json:"shadowEnable"`
				UID                string  `json:"uid"`
			} `json:"environment"`
			Ground struct {
				AoTextureUID  string    `json:"aoTextureUid"`
				Enable        bool      `json:"enable"`
				Fade          float64   `json:"fade"`
				Opacity       float64   `json:"opacity"`
				Position      []float64 `json:"position"`
				SamplingRange float64   `json:"samplingRange"`
				Scale         float64   `json:"scale"`
				ShadowMode    string    `json:"shadowMode"`
			} `json:"ground"`
			Hotspot struct {
				Hotspots []interface{} `json:"hotspots"`
				Visible  bool          `json:"visible"`
			} `json:"hotspot"`
			Lighting struct {
				Enable bool `json:"enable"`
				Lights []struct {
					Angle            int64     `json:"angle"`
					AttachedToCamera bool      `json:"attachedToCamera"`
					CastShadows      bool      `json:"castShadows"`
					Color            []float64 `json:"color"`
					Enable           bool      `json:"enable"`
					Falloff          int64     `json:"falloff"`
					Ground           []float64 `json:"ground"`
					Hardness         float64   `json:"hardness"`
					Intensity        float64   `json:"intensity"`
					IntensityGround  int64     `json:"intensityGround"`
					Matrix           []float64 `json:"matrix"`
					ShadowBias       float64   `json:"shadowBias"`
					Type             string    `json:"type"`
				} `json:"lights"`
			} `json:"lighting"`
			Materials struct {
				Nine301c684_327d_4a6d_bab1_0d9fe5a33295 struct {
					Channels struct {
						Aopbr struct {
							Enable          bool  `json:"enable"`
							Factor          int64 `json:"factor"`
							OccludeSpecular bool  `json:"occludeSpecular"`
						} `json:"AOPBR"`
						AlbedoPBR struct {
							Enable  bool  `json:"enable"`
							Factor  int64 `json:"factor"`
							Texture struct {
								InternalFormat string `json:"internalFormat"`
								MagFilter      string `json:"magFilter"`
								MinFilter      string `json:"minFilter"`
								TexCoordUnit   int64  `json:"texCoordUnit"`
								TextureTarget  string `json:"textureTarget"`
								UID            string `json:"uid"`
								WrapS          string `json:"wrapS"`
								WrapT          string `json:"wrapT"`
							} `json:"texture"`
						} `json:"AlbedoPBR"`
						Anisotropy struct {
							Direction int64 `json:"direction"`
							Enable    bool  `json:"enable"`
							Factor    int64 `json:"factor"`
							FlipXY    bool  `json:"flipXY"`
						} `json:"Anisotropy"`
						BumpMap struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"BumpMap"`
						CavityPBR struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"CavityPBR"`
						ClearCoat struct {
							Enable       bool      `json:"enable"`
							Factor       int64     `json:"factor"`
							Reflectivity int64     `json:"reflectivity"`
							Thickness    int64     `json:"thickness"`
							Tint         []float64 `json:"tint"`
						} `json:"ClearCoat"`
						ClearCoatNormalMap struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
							FlipY  bool  `json:"flipY"`
						} `json:"ClearCoatNormalMap"`
						ClearCoatRoughness struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"ClearCoatRoughness"`
						DiffuseColor struct {
							Color  []float64 `json:"color"`
							Enable bool      `json:"enable"`
							Factor float64   `json:"factor"`
						} `json:"DiffuseColor"`
						DiffuseIntensity struct {
							Color  []int64 `json:"color"`
							Enable bool    `json:"enable"`
							Factor int64   `json:"factor"`
						} `json:"DiffuseIntensity"`
						DiffusePBR struct {
							Color  []float64 `json:"color"`
							Enable bool      `json:"enable"`
							Factor int64     `json:"factor"`
						} `json:"DiffusePBR"`
						Displacement struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"Displacement"`
						EmitColor struct {
							Color  []float64 `json:"color"`
							Enable bool      `json:"enable"`
							Factor int64     `json:"factor"`
							Type   string    `json:"type"`
						} `json:"EmitColor"`
						GlossinessPBR struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"GlossinessPBR"`
						Matcap struct {
							Color     []int64 `json:"color"`
							Curvature int64   `json:"curvature"`
							Enable    bool    `json:"enable"`
							Factor    int64   `json:"factor"`
							Texture   struct {
								InternalFormat string `json:"internalFormat"`
								MagFilter      string `json:"magFilter"`
								MinFilter      string `json:"minFilter"`
								TexCoordUnit   int64  `json:"texCoordUnit"`
								TextureTarget  string `json:"textureTarget"`
								UID            string `json:"uid"`
								WrapS          string `json:"wrapS"`
								WrapT          string `json:"wrapT"`
							} `json:"texture"`
						} `json:"Matcap"`
						MetalnessPBR struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"MetalnessPBR"`
						NormalMap struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
							FlipY  bool  `json:"flipY"`
						} `json:"NormalMap"`
						Opacity struct {
							Enable                 bool    `json:"enable"`
							Factor                 int64   `json:"factor"`
							Invert                 bool    `json:"invert"`
							Ior                    float64 `json:"ior"`
							RefractionColor        []int64 `json:"refractionColor"`
							RoughnessFactor        int64   `json:"roughnessFactor"`
							ThinLayer              bool    `json:"thinLayer"`
							Type                   string  `json:"type"`
							UseAlbedoTint          bool    `json:"useAlbedoTint"`
							UseMicrosurfaceTexture bool    `json:"useMicrosurfaceTexture"`
							UseNormalOffset        bool    `json:"useNormalOffset"`
						} `json:"Opacity"`
						RoughnessPBR struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"RoughnessPBR"`
						SpecularColor struct {
							Color  []int64 `json:"color"`
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularColor"`
						SpecularF0 struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularF0"`
						SpecularHardness struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularHardness"`
						SpecularPBR struct {
							Color  []int64 `json:"color"`
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularPBR"`
						SubsurfaceScattering struct {
							Enable  bool    `json:"enable"`
							Factor  float64 `json:"factor"`
							Profile int64   `json:"profile"`
						} `json:"SubsurfaceScattering"`
						SubsurfaceTranslucency struct {
							Color           []int64 `json:"color"`
							Enable          bool    `json:"enable"`
							Factor          int64   `json:"factor"`
							ThicknessFactor float64 `json:"thicknessFactor"`
						} `json:"SubsurfaceTranslucency"`
					} `json:"channels"`
					CullFace   string  `json:"cullFace"`
					ID         string  `json:"id"`
					Name       string  `json:"name"`
					Reflection float64 `json:"reflection"`
					Shadeless  bool    `json:"shadeless"`
					StateSetID int64   `json:"stateSetID"`
					Version    int64   `json:"version"`
				} `json:"9301c684-327d-4a6d-bab1-0d9fe5a33295"`
				D245f255_a8d6_4f74_bab3_be8a903a4a58 struct {
					Channels struct {
						Aopbr struct {
							Enable          bool  `json:"enable"`
							Factor          int64 `json:"factor"`
							OccludeSpecular bool  `json:"occludeSpecular"`
						} `json:"AOPBR"`
						AlbedoPBR struct {
							Enable  bool  `json:"enable"`
							Factor  int64 `json:"factor"`
							Texture struct {
								InternalFormat string `json:"internalFormat"`
								MagFilter      string `json:"magFilter"`
								MinFilter      string `json:"minFilter"`
								TexCoordUnit   int64  `json:"texCoordUnit"`
								TextureTarget  string `json:"textureTarget"`
								UID            string `json:"uid"`
								WrapS          string `json:"wrapS"`
								WrapT          string `json:"wrapT"`
							} `json:"texture"`
						} `json:"AlbedoPBR"`
						Anisotropy struct {
							Direction int64 `json:"direction"`
							Enable    bool  `json:"enable"`
							Factor    int64 `json:"factor"`
							FlipXY    bool  `json:"flipXY"`
						} `json:"Anisotropy"`
						BumpMap struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"BumpMap"`
						CavityPBR struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"CavityPBR"`
						ClearCoat struct {
							Enable       bool      `json:"enable"`
							Factor       int64     `json:"factor"`
							Reflectivity int64     `json:"reflectivity"`
							Thickness    int64     `json:"thickness"`
							Tint         []float64 `json:"tint"`
						} `json:"ClearCoat"`
						ClearCoatNormalMap struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
							FlipY  bool  `json:"flipY"`
						} `json:"ClearCoatNormalMap"`
						ClearCoatRoughness struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"ClearCoatRoughness"`
						DiffuseColor struct {
							Color  []float64 `json:"color"`
							Enable bool      `json:"enable"`
							Factor float64   `json:"factor"`
						} `json:"DiffuseColor"`
						DiffuseIntensity struct {
							Color  []int64 `json:"color"`
							Enable bool    `json:"enable"`
							Factor int64   `json:"factor"`
						} `json:"DiffuseIntensity"`
						DiffusePBR struct {
							Color  []float64 `json:"color"`
							Enable bool      `json:"enable"`
							Factor int64     `json:"factor"`
						} `json:"DiffusePBR"`
						Displacement struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"Displacement"`
						EmitColor struct {
							Color  []float64 `json:"color"`
							Enable bool      `json:"enable"`
							Factor int64     `json:"factor"`
							Type   string    `json:"type"`
						} `json:"EmitColor"`
						GlossinessPBR struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"GlossinessPBR"`
						Matcap struct {
							Color     []int64 `json:"color"`
							Curvature int64   `json:"curvature"`
							Enable    bool    `json:"enable"`
							Factor    int64   `json:"factor"`
							Texture   struct {
								InternalFormat string `json:"internalFormat"`
								MagFilter      string `json:"magFilter"`
								MinFilter      string `json:"minFilter"`
								TexCoordUnit   int64  `json:"texCoordUnit"`
								TextureTarget  string `json:"textureTarget"`
								UID            string `json:"uid"`
								WrapS          string `json:"wrapS"`
								WrapT          string `json:"wrapT"`
							} `json:"texture"`
						} `json:"Matcap"`
						MetalnessPBR struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
						} `json:"MetalnessPBR"`
						NormalMap struct {
							Enable bool  `json:"enable"`
							Factor int64 `json:"factor"`
							FlipY  bool  `json:"flipY"`
						} `json:"NormalMap"`
						Opacity struct {
							Enable          bool    `json:"enable"`
							Factor          float64 `json:"factor"`
							Invert          bool    `json:"invert"`
							Ior             float64 `json:"ior"`
							RefractionColor []int64 `json:"refractionColor"`
							RoughnessFactor int64   `json:"roughnessFactor"`
							Texture         struct {
								InternalFormat string `json:"internalFormat"`
								MagFilter      string `json:"magFilter"`
								MinFilter      string `json:"minFilter"`
								TexCoordUnit   int64  `json:"texCoordUnit"`
								TextureTarget  string `json:"textureTarget"`
								UID            string `json:"uid"`
								WrapS          string `json:"wrapS"`
								WrapT          string `json:"wrapT"`
							} `json:"texture"`
							ThinLayer              bool   `json:"thinLayer"`
							Type                   string `json:"type"`
							UseAlbedoTint          bool   `json:"useAlbedoTint"`
							UseMicrosurfaceTexture bool   `json:"useMicrosurfaceTexture"`
							UseNormalOffset        bool   `json:"useNormalOffset"`
						} `json:"Opacity"`
						RoughnessPBR struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"RoughnessPBR"`
						SpecularColor struct {
							Color  []int64 `json:"color"`
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularColor"`
						SpecularF0 struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularF0"`
						SpecularHardness struct {
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularHardness"`
						SpecularPBR struct {
							Color  []int64 `json:"color"`
							Enable bool    `json:"enable"`
							Factor float64 `json:"factor"`
						} `json:"SpecularPBR"`
						SubsurfaceScattering struct {
							Enable  bool    `json:"enable"`
							Factor  float64 `json:"factor"`
							Profile int64   `json:"profile"`
						} `json:"SubsurfaceScattering"`
						SubsurfaceTranslucency struct {
							Color           []int64 `json:"color"`
							Enable          bool    `json:"enable"`
							Factor          int64   `json:"factor"`
							ThicknessFactor float64 `json:"thicknessFactor"`
						} `json:"SubsurfaceTranslucency"`
					} `json:"channels"`
					CullFace   string  `json:"cullFace"`
					ID         string  `json:"id"`
					Name       string  `json:"name"`
					Reflection float64 `json:"reflection"`
					Shadeless  bool    `json:"shadeless"`
					StateSetID int64   `json:"stateSetID"`
					Version    int64   `json:"version"`
				} `json:"d245f255-a8d6-4f74-bab3-be8a903a4a58"`
				UpdatedAt int64 `json:"updatedAt"`
			} `json:"materials"`
			Orientation struct {
				Matrix []int64 `json:"matrix"`
			} `json:"orientation"`
			Scene struct {
				PostProcess struct {
					Bloom struct {
						Enable    bool    `json:"enable"`
						Factor    float64 `json:"factor"`
						Radius    float64 `json:"radius"`
						Threshold int64   `json:"threshold"`
					} `json:"bloom"`
					ChromaticAberration struct {
						Enable bool    `json:"enable"`
						Factor float64 `json:"factor"`
					} `json:"chromaticAberration"`
					ColorBalance struct {
						Enable bool    `json:"enable"`
						High   []int64 `json:"high"`
						Low    []int64 `json:"low"`
						Mid    []int64 `json:"mid"`
					} `json:"colorBalance"`
					Dof struct {
						BlurFar    float64 `json:"blurFar"`
						BlurNear   float64 `json:"blurNear"`
						Enable     bool    `json:"enable"`
						FocusPoint []int64 `json:"focusPoint"`
					} `json:"dof"`
					Enable bool `json:"enable"`
					Grain  struct {
						Animated bool    `json:"animated"`
						Enable   bool    `json:"enable"`
						Factor   float64 `json:"factor"`
					} `json:"grain"`
					Sharpen struct {
						Enable bool    `json:"enable"`
						Factor float64 `json:"factor"`
					} `json:"sharpen"`
					Ssao struct {
						Bias      float64 `json:"bias"`
						Enable    bool    `json:"enable"`
						Intensity float64 `json:"intensity"`
						Radius    float64 `json:"radius"`
					} `json:"ssao"`
					Ssr struct {
						Enable bool  `json:"enable"`
						Factor int64 `json:"factor"`
					} `json:"ssr"`
					Taa struct {
						Enable      bool `json:"enable"`
						Transparent bool `json:"transparent"`
					} `json:"taa"`
					ToneMapping struct {
						Brightness int64  `json:"brightness"`
						Contrast   int64  `json:"contrast"`
						Enable     bool   `json:"enable"`
						Exposure   int64  `json:"exposure"`
						Method     string `json:"method"`
						Saturation int64  `json:"saturation"`
					} `json:"toneMapping"`
					Vignette struct {
						Amount   float64 `json:"amount"`
						Enable   bool    `json:"enable"`
						Hardness float64 `json:"hardness"`
					} `json:"vignette"`
				} `json:"postProcess"`
				SssProfiles []struct {
					Falloff  []int64   `json:"falloff"`
					Strength []float64 `json:"strength"`
				} `json:"sssProfiles"`
			} `json:"scene"`
			Shading struct {
				PointSize   int64  `json:"pointSize"`
				Renderer    string `json:"renderer"`
				Type        string `json:"type"`
				VertexColor struct {
					ColorSpace string `json:"colorSpace"`
					Enable     bool   `json:"enable"`
					UseAlpha   bool   `json:"useAlpha"`
				} `json:"vertexColor"`
			} `json:"shading"`
			Sound struct {
				Soundtracks []interface{} `json:"soundtracks"`
			} `json:"sound"`
			UpdatedAt string `json:"updatedAt"`
			Version   int64  `json:"version"`
			Vr        struct {
				DisplayFloor  bool    `json:"displayFloor"`
				FloorHeight   float64 `json:"floorHeight"`
				InitialCamera struct {
					Position []float64 `json:"position"`
					Rotation []int64   `json:"rotation"`
				} `json:"initialCamera"`
				WorldFactor float64 `json:"worldFactor"`
			} `json:"vr"`
			Wireframe struct {
				Color  string `json:"color"`
				Enable bool   `json:"enable"`
			} `json:"wireframe"`
		} `json:"options"`
		Org                     interface{} `json:"org"`
		PreferOriginalNormalMap bool        `json:"preferOriginalNormalMap"`
		ProcessedAt             string      `json:"processedAt"`
		ProcessingStatus        int64       `json:"processingStatus"`
		PublishedAt             string      `json:"publishedAt"`
		StaffpickedAt           string      `json:"staffpickedAt"`
		Status                  struct {
			Processing string   `json:"processing"`
			Warning    struct{} `json:"warning"`
		} `json:"status"`
		Thumbnails struct {
			Images []struct {
				Height int64  `json:"height"`
				Size   int64  `json:"size"`
				UID    string `json:"uid"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"images"`
			UID string `json:"uid"`
		} `json:"thumbnails"`
		UID  string `json:"uid"`
		User struct {
			Account string `json:"account"`
			Avatars struct {
				Images []struct {
					Height int64  `json:"height"`
					Size   int64  `json:"size"`
					URL    string `json:"url"`
					Width  int64  `json:"width"`
				} `json:"images"`
				UID string `json:"uid"`
			} `json:"avatars"`
			DisplayName string `json:"displayName"`
			ProfileURL  string `json:"profileUrl"`
			UID         string `json:"uid"`
			Username    string `json:"username"`
		} `json:"user"`
		Version struct {
			CreatedAt          string `json:"createdAt"`
			IsCandidateVersion bool   `json:"isCandidateVersion"`
			IsCurrentVersion   bool   `json:"isCurrentVersion"`
			ProcessedAt        string `json:"processedAt"`
			Reason             string `json:"reason"`
			Status             struct {
				Processing string   `json:"processing"`
				Warning    struct{} `json:"warning"`
			} `json:"status"`
			UID string `json:"uid"`
		} `json:"version"`
		VertexCount int64  `json:"vertexCount"`
		ViewCount   int64  `json:"viewCount"`
		ViewerURL   string `json:"viewerUrl"`
		Visibility  string `json:"visibility"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111"`
	I_models_29c76ea294264212b0598f358fbce111_animations_optimized_1 struct {
		Count   int64         `json:"count"`
		Results []interface{} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/animations?optimized=1"`
	I_models_29c76ea294264212b0598f358fbce111_backgrounds_optimized_1 struct {
		Count   int64 `json:"count"`
		Results []struct {
			CreatedAt string `json:"createdAt"`
			Images    []struct {
				CreatedAt string `json:"createdAt"`
				Height    int64  `json:"height"`
				Size      int64  `json:"size"`
				UpdatedAt string `json:"updatedAt"`
				URL       string `json:"url"`
				Width     int64  `json:"width"`
			} `json:"images"`
			IsDefault bool   `json:"isDefault"`
			Name      string `json:"name"`
			UID       string `json:"uid"`
			UpdatedAt string `json:"updatedAt"`
		} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/backgrounds?optimized=1"`
	I_models_29c76ea294264212b0598f358fbce111_environments_optimized_1 struct {
		Count   int64 `json:"count"`
		Results []struct {
			Brightness float64   `json:"brightness"`
			DiffuseSPH []float64 `json:"diffuseSPH"`
			IsDefault  bool      `json:"isDefault"`
			Lights     []struct {
				Area struct {
					H float64 `json:"h"`
					W float64 `json:"w"`
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"area"`
				Color      []float64 `json:"color"`
				Direction  []float64 `json:"direction"`
				Error      int64     `json:"error"`
				LumRatio   float64   `json:"lum_ratio"`
				Luminosity float64   `json:"luminosity"`
				Sum        float64   `json:"sum"`
				Variance   int64     `json:"variance"`
			} `json:"lights"`
			Name       string `json:"name"`
			Processing string `json:"processing"`
			Textures   []struct {
				Encoding string `json:"encoding"`
				Format   string `json:"format"`
				Images   []struct {
					Blur             float64 `json:"blur"`
					File             string  `json:"file"`
					Height           int64   `json:"height"`
					Name             string  `json:"name"`
					Samples          int64   `json:"samples"`
					SizeCompressed   int64   `json:"sizeCompressed"`
					SizeUncompressed int64   `json:"sizeUncompressed"`
					UID              string  `json:"uid"`
					Width            int64   `json:"width"`
				} `json:"images"`
				LimitSize int64  `json:"limitSize"`
				Type      string `json:"type"`
			} `json:"textures"`
			UID string `json:"uid"`
		} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/environments?optimized=1"`
	I_models_29c76ea294264212b0598f358fbce111_hotspots_optimized_1 struct {
		Count   int64         `json:"count"`
		Results []interface{} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/hotspots?optimized=1"`
	I_models_29c76ea294264212b0598f358fbce111_matcaps_optimized_1 struct {
		Count   int64 `json:"count"`
		Results []struct {
			CreatedAt string `json:"createdAt"`
			Images    []struct {
				CreatedAt string `json:"createdAt"`
				Height    int64  `json:"height"`
				Size      int64  `json:"size"`
				UpdatedAt string `json:"updatedAt"`
				URL       string `json:"url"`
				Width     int64  `json:"width"`
			} `json:"images"`
			IsDefault bool   `json:"isDefault"`
			Name      string `json:"name"`
			UID       string `json:"uid"`
			UpdatedAt string `json:"updatedAt"`
		} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/matcaps?optimized=1"`
	I_models_29c76ea294264212b0598f358fbce111_sounds_optimized_1 struct {
		Count   int64         `json:"count"`
		Results []interface{} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/sounds?optimized=1"`
	I_models_29c76ea294264212b0598f358fbce111_textures_optimized_1 struct {
		Count   int64 `json:"count"`
		Results []struct {
			ColorSpace string `json:"colorSpace"`
			CreatedAt  string `json:"createdAt"`
			Images     []struct {
				CreatedAt string `json:"createdAt"`
				Height    int64  `json:"height"`
				Options   struct {
					Format  string `json:"format"`
					Quality int64  `json:"quality"`
				} `json:"options"`
				Size      int64  `json:"size"`
				UID       string `json:"uid"`
				UpdatedAt string `json:"updatedAt"`
				URL       string `json:"url"`
				Width     int64  `json:"width"`
			} `json:"images"`
			Name      string `json:"name"`
			UID       string `json:"uid"`
			UpdatedAt string `json:"updatedAt"`
		} `json:"results"`
	} `json:"/i/models/29c76ea294264212b0598f358fbce111/textures?optimized=1"`
	I_users_me struct {
		AllowsRestricted bool     `json:"allowsRestricted"`
		Features         []string `json:"features"`
		IsAnonymous      bool     `json:"isAnonymous"`
	} `json:";/i/users/me"`
	DisplayStatus string `json:"displayStatus"`
	EmbedOptions  struct {
		AllowDownload                 bool        `json:"allowDownload"`
		AllowSwiftShader              bool        `json:"allowSwiftShader"`
		AnimationAutoplay             bool        `json:"animationAutoplay"`
		Anisotropy                    bool        `json:"anisotropy"`
		Annotation                    int64       `json:"annotation"`
		AnnotationCycle               interface{} `json:"annotationCycle"`
		AnnotationTooltipVisible      bool        `json:"annotationTooltipVisible"`
		AnnotationsVisible            interface{} `json:"annotationsVisible"`
		APIHookEnv                    interface{} `json:"apiHookEnv"`
		APILog                        interface{} `json:"apiLog"`
		APIVersion                    interface{} `json:"apiVersion"`
		Arkit                         bool        `json:"arkit"`
		ArkitDebug                    bool        `json:"arkitDebug"`
		AsyncImage                    bool        `json:"asyncImage"`
		AsyncShader                   int64       `json:"asyncShader"`
		AutoMaterial                  int64       `json:"autoMaterial"`
		Autospin                      int64       `json:"autospin"`
		Autostart                     bool        `json:"autostart"`
		Camera                        bool        `json:"camera"`
		CameraConstraints             bool        `json:"cameraConstraints"`
		CameraEasing                  interface{} `json:"cameraEasing"`
		CameraEye                     interface{} `json:"cameraEye"`
		CameraFollowBones             int64       `json:"cameraFollowBones"`
		CameraTarget                  interface{} `json:"cameraTarget"`
		Cardboard                     int64       `json:"cardboard"`
		Carmel                        bool        `json:"carmel"`
		CleanShader                   bool        `json:"cleanShader"`
		ContinuousRender              bool        `json:"continuousRender"`
		ConvertVertexColor8Bit        bool        `json:"convertVertexColor8Bit"`
		Debug3D                       bool        `json:"debug3D"`
		DepthMipmap                   int64       `json:"depthMipmap"`
		DisplayStatus                 string      `json:"displayStatus"`
		DofAttenuateDistance          bool        `json:"dofAttenuateDistance"`
		DofAttenuateSpeed             bool        `json:"dofAttenuateSpeed"`
		DofCircle                     bool        `json:"dofCircle"`
		DofRes                        int64       `json:"dofRes"`
		DofSticky                     bool        `json:"dofSticky"`
		DofTransition                 bool        `json:"dofTransition"`
		DoubleClick                   bool        `json:"doubleClick"`
		DownloadPicture               bool        `json:"downloadPicture"`
		Drs                           bool        `json:"drs"`
		DrsFps                        int64       `json:"drsFps"`
		DrsMin                        interface{} `json:"drsMin"`
		DrsRatio                      int64       `json:"drsRatio"`
		DrsTest                       bool        `json:"drsTest"`
		EpsilonAlpha                  int64       `json:"epsilonAlpha"`
		FloatRtt                      int64       `json:"floatRtt"`
		ForceController               string      `json:"forceController"`
		ForceControllerOrientation    float64     `json:"forceControllerOrientation"`
		ForceControllerRayOrientation float64     `json:"forceControllerRayOrientation"`
		ForceFallback                 bool        `json:"forceFallback"`
		ForceMs                       int64       `json:"forceMs"`
		ForceTriangles                interface{} `json:"forceTriangles"`
		FpsSpeed                      int64       `json:"fpsSpeed"`
		Fxaa                          interface{} `json:"fxaa"`
		GrainSpeed                    int64       `json:"grainSpeed"`
		GraphOptimizer                interface{} `json:"graphOptimizer"`
		ImageCompression              interface{} `json:"imageCompression"`
		InStore                       bool        `json:"inStore"`
		Internal                      bool        `json:"internal"`
		IsApp                         bool        `json:"isApp"`
		KeepEmptyGeometries           bool        `json:"keepEmptyGeometries"`
		MaterialAo                    interface{} `json:"materialAo"`
		MaterialCavity                interface{} `json:"materialCavity"`
		MaterialDiffuse               interface{} `json:"materialDiffuse"`
		MaterialDisplacement          interface{} `json:"materialDisplacement"`
		MaterialDisplacementFactor    interface{} `json:"materialDisplacementFactor"`
		MaterialEmissive              interface{} `json:"materialEmissive"`
		MaterialF0                    interface{} `json:"materialF0"`
		MaterialGlossiness            interface{} `json:"materialGlossiness"`
		MaterialMetalness             interface{} `json:"materialMetalness"`
		MaterialNames                 string      `json:"materialNames"`
		MaterialNormal                interface{} `json:"materialNormal"`
		MaterialPacking               interface{} `json:"materialPacking"`
		MaterialRoughness             interface{} `json:"materialRoughness"`
		MaterialShowcase              bool        `json:"materialShowcase"`
		MaterialSpecular              interface{} `json:"materialSpecular"`
		MaterialTransparency          interface{} `json:"materialTransparency"`
		MaxDevicePixelRatio           float64     `json:"maxDevicePixelRatio"`
		MaxTextureSize                int64       `json:"maxTextureSize"`
		MaxTextureUnits               int64       `json:"maxTextureUnits"`
		MaxVertexUniforms             int64       `json:"maxVertexUniforms"`
		MergeMaterials                interface{} `json:"mergeMaterials"`
		Messaging                     interface{} `json:"messaging"`
		Model                         interface{} `json:"model"`
		ModelOptimization             interface{} `json:"modelOptimization"`
		MorphEpsilon                  float64     `json:"morphEpsilon"`
		MorphGpu                      int64       `json:"morphGpu"`
		Navigation                    string      `json:"navigation"`
		OrbitConstraintPan            interface{} `json:"orbitConstraintPan"`
		OrbitConstraintPitchDown      interface{} `json:"orbitConstraintPitchDown"`
		OrbitConstraintPitchUp        interface{} `json:"orbitConstraintPitchUp"`
		OrbitConstraintYawLeft        interface{} `json:"orbitConstraintYawLeft"`
		OrbitConstraintYawRight       interface{} `json:"orbitConstraintYawRight"`
		OrbitConstraintZoomIn         interface{} `json:"orbitConstraintZoomIn"`
		OrbitConstraintZoomOut        interface{} `json:"orbitConstraintZoomOut"`
		OrbitPanFactor                int64       `json:"orbitPanFactor"`
		OrbitRotationFactor           int64       `json:"orbitRotationFactor"`
		OrbitZoomFactor               int64       `json:"orbitZoomFactor"`
		Panorama                      bool        `json:"panorama"`
		PixelBudget                   interface{} `json:"pixelBudget"`
		PostProcess                   interface{} `json:"postProcess"`
		PostProcessASCII              interface{} `json:"postProcessAscii"`
		PowerPreference               string      `json:"powerPreference"`
		Pratio                        interface{} `json:"pratio"`
		Preload                       bool        `json:"preload"`
		PreserveDrawingBuffer         bool        `json:"preserveDrawingBuffer"`
		ProcessMaterial               bool        `json:"processMaterial"`
		ProcessOptions                bool        `json:"processOptions"`
		Quality                       interface{} `json:"quality"`
		RefractionRes                 int64       `json:"refractionRes"`
		ResizeTimeout                 int64       `json:"resizeTimeout"`
		Rgbm                          bool        `json:"rgbm"`
		Scale                         int64       `json:"scale"`
		ScaleEpsilon                  int64       `json:"scaleEpsilon"`
		Scrollwheel                   bool        `json:"scrollwheel"`
		Shadow                        bool        `json:"shadow"`
		ShadowAtlas                   bool        `json:"shadowAtlas"`
		ShadowJitterOffset            string      `json:"shadowJitterOffset"`
		ShadowNormalOffset            bool        `json:"shadowNormalOffset"`
		ShadowPCF                     string      `json:"shadowPCF"`
		ShadowTextureSize             int64       `json:"shadowTextureSize"`
		Share                         bool        `json:"share"`
		SingleSided                   interface{} `json:"singleSided"`
		SnapKeyframe                  bool        `json:"snapKeyframe"`
		SoundEnable                   bool        `json:"soundEnable"`
		SoundMute                     bool        `json:"soundMute"`
		SoundPreload                  bool        `json:"soundPreload"`
		SplitLimit                    int64       `json:"splitLimit"`
		SsaoNormal                    bool        `json:"ssaoNormal"`
		Ssr                           interface{} `json:"ssr"`
		SsrTransparent                bool        `json:"ssrTransparent"`
		SssJitter                     int64       `json:"sssJitter"`
		SssKernel                     int64       `json:"sssKernel"`
		Stats                         bool        `json:"stats"`
		SubstanceDisplacement         bool        `json:"substanceDisplacement"`
		Supersample                   bool        `json:"supersample"`
		SvgSize                       int64       `json:"svgSize"`
		Taa                           interface{} `json:"taa"`
		TaaAnimation                  bool        `json:"taaAnimation"`
		TaaFeedbackMax                float64     `json:"taaFeedbackMax"`
		TaaFeedbackMin                float64     `json:"taaFeedbackMin"`
		TaaJitter                     interface{} `json:"taaJitter"`
		TaaTransparent                interface{} `json:"taaTransparent"`
		Tab                           interface{} `json:"tab"`
		TextureFrameBudget            int64       `json:"textureFrameBudget"`
		Tracking                      bool        `json:"tracking"`
		Transparent                   bool        `json:"transparent"`
		UIAnimations                  bool        `json:"uiAnimations"`
		UIAnnotations                 bool        `json:"uiAnnotations"`
		UIAr                          bool        `json:"uiAr"`
		UIArHelp                      bool        `json:"uiArHelp"`
		UIArQrcode                    bool        `json:"uiArQrcode"`
		UIColor                       string      `json:"uiColor"`
		UIControls                    bool        `json:"uiControls"`
		UIFadeout                     bool        `json:"uiFadeout"`
		UIFullscreen                  bool        `json:"uiFullscreen"`
		UIGeneralControls             bool        `json:"uiGeneralControls"`
		UIHelp                        bool        `json:"uiHelp"`
		UIHint                        int64       `json:"uiHint"`
		UIInfos                       bool        `json:"uiInfos"`
		UIInspector                   bool        `json:"uiInspector"`
		UIInspectorOpen               bool        `json:"uiInspectorOpen"`
		UILoading                     bool        `json:"uiLoading"`
		UISettings                    bool        `json:"uiSettings"`
		UISnapshots                   bool        `json:"uiSnapshots"`
		UISound                       bool        `json:"uiSound"`
		UIStart                       bool        `json:"uiStart"`
		UIStop                        bool        `json:"uiStop"`
		UITheatre                     bool        `json:"uiTheatre"`
		UITheme                       interface{} `json:"uiTheme"`
		UIVr                          bool        `json:"uiVr"`
		UIWatermark                   bool        `json:"uiWatermark"`
		UIWatermarkLink               bool        `json:"uiWatermarkLink"`
		UnitRenderLocal               bool        `json:"unitRenderLocal"`
		UnitRenderReporter            string      `json:"unitRenderReporter"`
		UnitRenderTask                interface{} `json:"unitRenderTask"`
		UseVao                        bool        `json:"useVao"`
		VaryingSorting                bool        `json:"varyingSorting"`
		VrAr                          bool        `json:"vrAr"`
		VrBrowserEnv                  string      `json:"vrBrowserEnv"`
		VrFade                        int64       `json:"vrFade"`
		VrForceRaf                    bool        `json:"vrForceRaf"`
		VrGenerateUITextures          bool        `json:"vrGenerateUiTextures"`
		VrInNavigation                int64       `json:"vrInNavigation"`
		VrLauncherAlphaCardFactor     int64       `json:"vrLauncherAlphaCardFactor"`
		VrLinkNavigation              int64       `json:"vrLinkNavigation"`
		VrMaxRttSize                  interface{} `json:"vrMaxRttSize"`
		VrMirror                      bool        `json:"vrMirror"`
		VrMobileMaxFaces              int64       `json:"vrMobileMaxFaces"`
		VrQuality                     interface{} `json:"vrQuality"`
		VrScaleRtt                    int64       `json:"vrScaleRtt"`
		VrStereo                      bool        `json:"vrStereo"`
		Webgl2                        interface{} `json:"webgl2"`
		WebglExt                      bool        `json:"webglExt"`
		WebglRenderTexture            interface{} `json:"webglRenderTexture"`
		WebglRestore                  int64       `json:"webglRestore"`
		WebglTimerGpu                 bool        `json:"webglTimerGpu"`
		WebglUniforms                 bool        `json:"webglUniforms"`
		WireframePreload              interface{} `json:"wireframePreload"`
		Zoct                          bool        `json:"zoct"`
		Zq                            bool        `json:"zq"`
		Zratio                        int64       `json:"zratio"`
		Zw                            bool        `json:"zw"`
		Zz                            interface{} `json:"zz"`
	} `json:"embedOptions"`
}
